package hiker

// Ginkgo does not find your specs by itself. RunSpecs below is what hands
// them to go test, so this file has to exist and has to call it. Without it
// everything still compiles and no spec runs at all.
//
// There is one of these per suite, not one per spec file. Adding more spec
// files needs nothing here.

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHiker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hiker Suite")
}
