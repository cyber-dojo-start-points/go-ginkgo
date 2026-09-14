package hiker

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("answer", func() {

	It("is life the universe and everything", func() {
		Expect(answer()).To(Equal(42))
	})

	It("is even", func() {
		Expect(answer() % 2).To(Equal(0))
	})
})
