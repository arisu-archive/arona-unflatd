package cmd_test

import (
	"errors"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd"
)

var _ = Describe("Command", func() {
	var (
		logger  *slog.Logger
		testCmd *cobra.Command
	)

	BeforeEach(func() {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		testCmd = &cobra.Command{Use: "test"}
	})

	Describe("RunE", func() {
		Context("when the command executes successfully", func() {
			It("should log success and return nil", func() {
				executor := cmd.RunE("test-verb", logger, func(*cobra.Command, []string) error {
					return nil
				})

				err := executor(testCmd, []string{})
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the command returns an error", func() {
			It("should wrap the error with timing information", func() {
				expectedErr := errors.New("command failed")
				executor := cmd.RunE("test-verb", logger, func(*cobra.Command, []string) error {
					return expectedErr
				})

				err := executor(testCmd, []string{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("test-verb failed after"))
				Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
			})
		})

		Context("when the command panics", func() {
			It("should recover and return an error", func() {
				executor := cmd.RunE("test-verb", logger, func(*cobra.Command, []string) error {
					panic("unexpected panic")
				})

				err := executor(testCmd, []string{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("test-verb failed after"))
				Expect(err.Error()).To(ContainSubstring("unexpected panic"))
			})
		})
	})
})
