package unflatd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/spf13/cobra"

	"github.com/arisu-archive/arona-unflatd/cmd"
	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/conversion"
	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/query"
)

type Command struct {
	cmd       *cobra.Command
	converter *conversion.SchemaConverter
	logger    *slog.Logger
	opts      Options
}

func NewCommand(logger *slog.Logger) *Command {
	unflatd := &Command{
		converter: conversion.NewSchemaConverter(logger),
		logger:    logger,
	}
	cobraCmd := &cobra.Command{
		Use:               "decompile",
		Aliases:           []string{"d"},
		Short:             "Generate FlatBuffer schema from decompiled C# code",
		SilenceErrors:     true,
		SilenceUsage:      true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE:              cmd.RunE("decompile", logger, unflatd.Execute),
	}
	cobraCmd.Flags().StringVarP(&unflatd.opts.input, "input", "i", "", "Input directory containing decompiled C# code")
	cobraCmd.Flags().StringVarP(
		&unflatd.opts.output,
		"output",
		"o",
		"",
		"Output directory for the generated FlatBuffer schema",
	)
	cobraCmd.Flags().StringVarP(
		&unflatd.opts.namespace,
		"namespace",
		"n",
		"FlatData",
		"Filter out the mismatched namespace in the generated FlatBuffer schema",
	)
	if err := cobraCmd.MarkFlagRequired("input"); err != nil {
		panic("failed to mark input flag as required: " + err.Error())
	}
	if err := cobraCmd.MarkFlagRequired("output"); err != nil {
		panic("failed to mark output flag as required: " + err.Error())
	}

	unflatd.cmd = cobraCmd
	return unflatd
}

func (c *Command) Command() *cobra.Command {
	return c.cmd
}

type ProcessedFile struct {
	Path   string
	Schema *fbs.Schema
}

func (c *Command) Execute(_ *cobra.Command, _ []string) error {
	c.logger.Debug("Decompiling C# code", "input", c.opts.input, "output", c.opts.output, "namespace", c.opts.namespace)
	// Create output directory if not exists
	if err := os.MkdirAll(c.opts.output, 0o700); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	files, err := doublestar.Glob(os.DirFS(c.opts.input), "**/*.cs")
	if err != nil {
		return fmt.Errorf("failed to glob input directory: %w", err)
	}
	c.logger.Debug("Found files", "count", len(files))

	parser, err := query.NewParser()
	if err != nil {
		return fmt.Errorf("failed to create parser: %w", err)
	}
	processedFiles := make([]ProcessedFile, 0)
	for _, file := range files {
		fullPath := filepath.Join(c.opts.input, file)
		c.logger.Info("Processing file", "file", fullPath)
		structs, parseErr := parser.ProcessFile(context.Background(), fullPath)
		if parseErr != nil {
			return fmt.Errorf("failed to process file: %w", parseErr)
		}
		schema, conversionErr := c.converter.Convert(structs)
		if conversionErr != nil {
			if errors.Is(conversionErr, conversion.ErrNoTablesOrEnumsFound) {
				c.logger.Debug("No tables or enums found in file", "file", fullPath)
				continue
			}
			return fmt.Errorf("failed to convert flatbuffer: %w", conversionErr)
		}
		c.logger.Debug("FlatBuffer schema", "fbs", schema)
		c.logger.Info("Generated FlatBuffer schema", "file", fullPath)
		// Mark the reference count of each schema.
		processedFiles = append(processedFiles, ProcessedFile{Path: fullPath, Schema: schema})
	}

	// Post processing: Fixing the imports
	fixImports(processedFiles)
	refCount := make(fbs.SchemaReference)
	for _, file := range processedFiles {
		for _, imp := range file.Schema.Imports {
			refCount[imp] = append(refCount[imp], file.Schema)
		}
	}

	// Remove the schemas that are not referenced and not equal to the namespace.
	var filteredFiles []ProcessedFile
	for _, file := range processedFiles {
		// No reference and not equal to the namespace
		schemaFile := strings.TrimSuffix(filepath.Base(file.Path), filepath.Ext(file.Path))
		if !refCount.HasNamespace(schemaFile, c.opts.namespace) && file.Schema.Namespace != c.opts.namespace {
			c.logger.Debug("Removing dangling schema", "path", file.Path, "namespace", file.Schema.Namespace)
			continue
		}

		c.logger.Debug("This is a valid schema", "path", file.Path, "namespace", file.Schema.Namespace)
		// Force the namespace to the specified namespace (As go flatc cannot handle flatbuffers without namespace)
		file.Schema.Namespace = c.opts.namespace
		filteredFiles = append(filteredFiles, file)
	}

	// Write all the schemas to output
	for _, file := range filteredFiles {
		baseFile := strings.TrimSuffix(filepath.Base(file.Path), filepath.Ext(file.Path))
		outputPath := filepath.Join(c.opts.output, baseFile+".fbs")
		c.logger.Debug("Schema", "schema", file.Schema)
		c.logger.Info("Writing FlatBuffer schema", "file", outputPath)
		v := fbs.NewSchemaVisitor()
		result := v.VisitSchema(file.Schema)
		c.logger.Debug("Generated FlatBuffer source code", "result", result)
		if writeErr := os.WriteFile(outputPath, []byte(result), 0o600); writeErr != nil {
			return fmt.Errorf("failed to write FlatBuffer schema: %w", writeErr)
		}
	}
	return nil
}

func fixImports(files []ProcessedFile) {
	for _, file := range files {
		imports := make(map[string]struct{})
		for _, table := range file.Schema.Tables {
			for _, field := range table.Fields {
				if field.Namespace != "" {
					field.Type = strings.TrimPrefix(field.Type, field.Namespace+".")
				}
				if !field.IsPrimitive() {
					imports[field.Type] = struct{}{}
				}
			}
		}

		// Convert unique imports to slice
		file.Schema.Imports = make([]string, 0, len(imports))
		for imp := range imports {
			file.Schema.Imports = append(file.Schema.Imports, imp)
		}
	}
}
