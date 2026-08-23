package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version metadata", func() {
	DescribeTable("reports the build version",
		func(ldflags, want string) {
			binaryName := "arona-unflatd"
			if runtime.GOOS == "windows" {
				binaryName += ".exe"
			}
			binaryPath := filepath.Join(GinkgoT().TempDir(), binaryName)

			buildArgs := []string{"build"}
			if ldflags != "" {
				buildArgs = append(buildArgs, "-ldflags="+ldflags)
			}
			buildArgs = append(buildArgs, "-o", binaryPath, ".")

			build := exec.Command("go", buildArgs...)
			buildOutput, err := build.CombinedOutput()
			Expect(err).NotTo(HaveOccurred(), string(buildOutput))

			command := exec.Command(binaryPath, "--version")
			commandOutput, err := command.CombinedOutput()
			Expect(err).NotTo(HaveOccurred(), string(commandOutput))
			Expect(strings.TrimSpace(string(commandOutput))).To(Equal("arona-unflatd version " + want))
		},
		Entry("for source builds", "", "dev"),
		Entry("injected by the linker", "-X main.version=9.8.7", "9.8.7"),
	)
})
