package hiker

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// FIt focuses the first spec, so ginkgo skips the second one, which asserts
// three digits and would otherwise fail. Ginkgo refuses to let that pass as a
// clean run: it exits 197 and go test reports FAIL, so the honest colour is
// amber.
var _ = Describe("answer", func() {

	FIt("is life the universe and everything", func() {
		Expect(answer()).To(Equal(42))
	})

	It("is three digits", func() {
		Expect(answer()).To(BeNumerically(">=", 100))
	})
})
