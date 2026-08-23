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
	It("reports the version injected by the linker", func() {
		binaryName := "arona-unflatd"
		if runtime.GOOS == "windows" {
			binaryName += ".exe"
		}
		binaryPath := filepath.Join(GinkgoT().TempDir(), binaryName)

		build := exec.Command(
			"go",
			"build",
			"-ldflags=-X main.version=9.8.7",
			"-o",
			binaryPath,
			".",
		)
		buildOutput, err := build.CombinedOutput()
		Expect(err).NotTo(HaveOccurred(), string(buildOutput))

		command := exec.Command(binaryPath, "--version")
		commandOutput, err := command.CombinedOutput()
		Expect(err).NotTo(HaveOccurred(), string(commandOutput))
		Expect(strings.TrimSpace(string(commandOutput))).To(Equal("arona-unflatd version 9.8.7"))
	})
})
