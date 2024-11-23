package unflatd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUnflatd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unflatd Suite")
}
