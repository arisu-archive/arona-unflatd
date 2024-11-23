package fbs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFbs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FBS Suite")
}
