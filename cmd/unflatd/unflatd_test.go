package unflatd_test

import (
	"log/slog"
	"os"
	"path/filepath"

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

	Context("when input directory doesn't exist", func() {
		It("should return an error", func() {
			cmd := unflatd.NewCommand(logger)
			cobraCmd := cmd.Command()
			cobraCmd.SetArgs([]string{
				"--input", "/nonexistent/path",
				"--output", outputDir,
			})

			err := cmd.Execute(cmd.Command(), []string{})
			Expect(err).To(HaveOccurred())
		})
	})
})
