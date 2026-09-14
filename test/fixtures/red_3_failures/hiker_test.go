package hiker

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("answer", func() {

	It("is life the universe and everything", func() {
		Expect(answer()).To(Equal(42))
	})

	It("is less than fifty", func() {
		Expect(answer()).To(BeNumerically("<", 50))
	})

	It("is a multiple of seven", func() {
		Expect(answer() % 7).To(Equal(0))
	})
})
