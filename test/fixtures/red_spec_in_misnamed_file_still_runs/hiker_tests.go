package hiker

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// This file ends _tests.go, so go test does not collect test functions from
// it. A ginkgo spec does not depend on that: Describe runs at package init and
// registers the spec with the global suite, and the file is compiled into the
// package either way. So this spec runs, and it asserts three digits, which
// takes the whole suite red.
var _ = Describe("the size of the answer", func() {

	It("is three digits", func() {
		Expect(answer()).To(BeNumerically(">=", 100))
	})
})
