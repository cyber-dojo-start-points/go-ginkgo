package hiker

// The suite bootstrap is here and RunSpecs is called, but no file registers a
// spec, so ginkgo has nothing to run. It calls that a success, which is why
// the count of specs rather than the verdict is what keeps this out of green.

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHiker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hiker Suite")
}
