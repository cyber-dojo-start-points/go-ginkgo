package hiker

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("the size of the answer", func() {

	It("is two digits", func() {
		Expect(answer()).To(BeNumerically(">=", 10))
		Expect(answer()).To(BeNumerically("<=", 99))
	})

	It("is three digits", func() {
		Expect(answer()).To(BeNumerically(">=", 100))
	})
})
