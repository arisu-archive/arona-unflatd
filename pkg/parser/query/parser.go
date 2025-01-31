package query

import (
	"context"
	"fmt"
	"os"

	"github.com/smacker/go-tree-sitter/csharp"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/arisu-archive/arona-unflatd/pkg/parser/analyzer"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/state"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/visitor"
)

type Parser struct {
	query *sitter.Query
	v     *visitor.ASTVisitor
}

func NewParser() (*Parser, error) {
	directQuery, err := sitter.NewQuery([]byte(StructParsingQuery), csharp.GetLanguage())
	if err != nil {
		return nil, fmt.Errorf("failed to create direct query: %w", err)
	}

	return &Parser{query: directQuery, v: visitor.NewASTVisitor()}, nil
}

func (p *Parser) ProcessFile(ctx context.Context, path string) (*ast.FileInfo, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	tree, err := sitter.ParseCtx(ctx, content, csharp.GetLanguage())
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}

	return p.processTree(tree, content)
}

func (p *Parser) processTree(rootNode *sitter.Node, content []byte) (*ast.FileInfo, error) {
	qc := sitter.NewQueryCursor()
	qc.Exec(p.query, rootNode)

	for {
		ps := &analyzer.ParsingState{
			Enum: ast.EnumInfo{
				BaseType: "int",
			},
		}
		match, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, capture := range match.Captures {
			p.processCapture(capture, content, ps)
		}

		// Determine and execute the appropriate state
		nodeState := state.DetermineState(ps)
		nodeState.Visit(p.v, ps)
		ps.Reset()
	}

	return p.v.Result(), nil
}

func (p *Parser) processCapture(capture sitter.QueryCapture, content []byte, ps *analyzer.ParsingState) {
	captureName := p.query.CaptureNameForId(capture.Index)
	nodeContent := capture.Node.Content(content)

	switch captureName {
	case "namespace":
		ps.HandleNamespace(nodeContent)
	case "struct_name":
		ps.HandleName(nodeContent)
	case "modifier", "method_modifier", "enum_modifier", "field_modifier":
		ps.HandleModifier(nodeContent)
	case "field_name":
		ps.HandleFieldName(nodeContent)
	case "field_type":
		ps.HandleFieldType(nodeContent)
	case "method_name":
		ps.HandleMethodName(nodeContent)
	case "method_return_type":
		ps.HandleMethodReturnType(nodeContent)
	case "param_name":
		ps.HandleParamName(nodeContent)
	case "param_type":
		ps.HandleParamType(nodeContent)
	case "enum_name":
		ps.HandleEnumName(nodeContent)
	case "enum_base_type":
		ps.HandleEnumBaseType(nodeContent)
	case "enum_member_name":
		ps.HandleEnumMemberName(nodeContent)
	case "enum_member_value":
		ps.HandleEnumMemberValue(nodeContent)
	case "interface":
		ps.HandleInterface(nodeContent)
	}
}
