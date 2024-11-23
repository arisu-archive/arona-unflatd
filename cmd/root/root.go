package root

import (
	"io"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd"
)

type ExecuteOptions struct {
	Version string
	Exit    func(int)
	In      io.Reader
	Out     io.Writer
	Err     io.Writer
}

type rootCmd struct {
	cmd      *cobra.Command
	exit     func(int)
	verbose  bool
	logger   *slog.Logger
	logLevel *slog.LevelVar
}

func Execute(opts ExecuteOptions, args []string) {
	newRootCmd(opts).Execute(args)
}

func (r *rootCmd) Execute(args []string) {
	r.cmd.SetArgs(args)
	if err := r.cmd.Execute(); err != nil {
		r.logger.Error("arona-unflatd failed.", slog.Any("error", err))
		r.exit(1)
	}
}

func newRootCmd(opts ExecuteOptions) *rootCmd {
	logLevel := new(slog.LevelVar)
	logLevel.Set(slog.LevelInfo)

	root := &rootCmd{
		exit:     opts.Exit,
		logger:   slog.New(slog.NewTextHandler(opts.Out, &slog.HandlerOptions{Level: logLevel})),
		logLevel: logLevel,
	}

	cmd := &cobra.Command{
		Use:   "arona-unflatd <command> [flags]",
		Short: "Arona Unflatd - FlatBuffer schema generator",
		Long: `AronaUnflatd is a specialized tool designed to reconstruct FlatBuffer schema files
from decompiled C# code. It analyzes the structure and attributes of decompiled
FlatBuffer-generated C# classes and automatically generates the corresponding .fbs
schema files, making it easier to recover original schemas when they're not available.`,
		Version:           opts.Version,
		SilenceErrors:     false,
		SilenceUsage:      false,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if root.verbose {
				root.logLevel.Set(slog.LevelDebug)
				root.logger.Debug("verbose mode enabled")
			}
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			root.logger.Info("arona-unflatd completed successfully")
		},
	}

	cmd.PersistentFlags().BoolVarP(&root.verbose, "verbose", "v", false, "Enable verbose mode")
	cmd.AddCommand(unflatd.NewCommand(root.logger).Command())
	cmd.SetIn(opts.In)
	cmd.SetOut(opts.Out)
	cmd.SetErr(opts.Err)

	root.cmd = cmd
	return root
}
