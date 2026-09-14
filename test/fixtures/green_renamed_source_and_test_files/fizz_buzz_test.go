package fizzbuzz

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("fizzBuzz", func() {

	It("says FizzBuzz for a multiple of three and five", func() {
		Expect(fizzBuzz(15)).To(Equal("FizzBuzz"))
	})
})
