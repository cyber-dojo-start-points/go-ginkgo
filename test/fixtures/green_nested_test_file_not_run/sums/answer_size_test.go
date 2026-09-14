package sums

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// The print is what shows whether this spec ran, and the expectation asserts
// three digits, so the case would go red as well as printing if it did.
var _ = Describe("the size of the answer", func() {

	It("is three digits", func() {
		fmt.Println("checking the size of the answer")
		Expect(42).To(BeNumerically(">=", 100))
	})
})
