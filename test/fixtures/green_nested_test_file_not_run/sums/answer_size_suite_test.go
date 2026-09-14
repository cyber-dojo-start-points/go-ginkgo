package sums

// Go ties a package to its directory, and cyber-dojo.sh runs go test without
// ./..., so go test only ever builds the sandbox directory. This sub-directory
// is a package of its own and nothing here is compiled, not even this suite
// bootstrap, so the specs beside it are never handed to ginkgo.

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSums(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sums Suite")
}
