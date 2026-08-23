package unflatd_test

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd"
)

var _ = Describe("Unflatd", func() {
	var (
		tmpDir    string
		inputDir  string
		outputDir string
		logger    *slog.Logger
		cmd       *unflatd.Command
	)

	BeforeEach(func() {
		var err error
		tmpDir, err = os.MkdirTemp("", "unflatd-test-*")
		Expect(err).NotTo(HaveOccurred())

		inputDir = filepath.Join(tmpDir, "input")
		outputDir = filepath.Join(tmpDir, "output")

		err = os.MkdirAll(inputDir, 0o755)
		Expect(err).NotTo(HaveOccurred())

		// Setup test logger
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		cmd = unflatd.NewCommand(logger)

		// Set required flags via cobra command
		cobraCmd := cmd.Command()
		cobraCmd.SetContext(context.Background())
		err = cobraCmd.Flags().Set("input", inputDir)
		Expect(err).NotTo(HaveOccurred())
		err = cobraCmd.Flags().Set("output", outputDir)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tmpDir)
	})

	Context("when processing valid C# files", func() {
		BeforeEach(func() {
			// Create a test C# file
			testFile := filepath.Join(inputDir, "TestEnum.cs")
			content := `
namespace FlatData;

public enum TestEnum
{
	None = 0,
	First = 1,
	Second = 2
}
`
			err := os.WriteFile(testFile, []byte(content), 0o644)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should generate flatbuffer schema files", func() {
			err := cmd.Execute(cmd.Command(), []string{})
			Expect(err).NotTo(HaveOccurred())

			// Verify output file exists
			outputFile := filepath.Join(outputDir, "TestEnum.fbs")
			Expect(outputFile).To(BeARegularFile())

			// Verify content
			content, err := os.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			schemaContent := string(content)
			Expect(schemaContent).To(ContainSubstring("namespace FlatData;"))
			Expect(schemaContent).To(ContainSubstring("enum TestEnum"))
			Expect(schemaContent).To(ContainSubstring("None = 0"))
			Expect(schemaContent).To(ContainSubstring("First = 1"))
			Expect(schemaContent).To(ContainSubstring("Second = 2"))
		})
	})

	Context("when recovering a generated FlatBuffer type", func() {
		It("should match the golden schema", func() {
			goldenCommand := unflatd.NewCommand(logger)
			cobraCmd := goldenCommand.Command()
			cobraCmd.SetContext(context.Background())
			err := cobraCmd.Flags().Set("input", filepath.Join("testdata", "recovery", "input"))
			Expect(err).NotTo(HaveOccurred())
			err = cobraCmd.Flags().Set("output", outputDir)
			Expect(err).NotTo(HaveOccurred())

			err = goldenCommand.Execute(cobraCmd, nil)
			Expect(err).NotTo(HaveOccurred())

			got, err := os.ReadFile(filepath.Join(outputDir, "TestSchema.fbs"))
			Expect(err).NotTo(HaveOccurred())
			want, err := os.ReadFile(filepath.Join("testdata", "recovery", "want", "TestSchema.fbs"))
			Expect(err).NotTo(HaveOccurred())

			normalize := func(content []byte) string {
				return strings.TrimRight(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
			}
			Expect(normalize(got)).To(Equal(normalize(want)))
		})
	})

	Context("when input directory doesn't exist", func() {
		It("should return an error", func() {
			unflatCmd := unflatd.NewCommand(logger)
			cobraCmd := unflatCmd.Command()
			cobraCmd.SetArgs([]string{
				"--input", "/nonexistent/path",
				"--output", outputDir,
			})

			err := unflatCmd.Execute(unflatCmd.Command(), []string{})
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when command context is canceled", func() {
		It("should stop C# parsing", func() {
			testFile := filepath.Join(inputDir, "Canceled.cs")
			err := os.WriteFile(testFile, []byte("namespace FlatData; public enum Canceled { None = 0 }"), 0o600)
			Expect(err).NotTo(HaveOccurred())

			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			cobraCmd := cmd.Command()
			cobraCmd.SetContext(ctx)

			err = cmd.Execute(cobraCmd, nil)

			Expect(err).To(MatchError(ContainSubstring("context canceled")))
		})
	})
})
