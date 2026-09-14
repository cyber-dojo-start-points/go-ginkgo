package hiker

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Ginkgo prints a single bullet for a spec that passes and names no file, so
// this print is what shows this file was compiled into the suite and run.
var _ = Describe("the size of the answer", func() {

	It("is two digits", func() {
		fmt.Println("checking the size of the answer")
		Expect(answer()).To(BeNumerically(">=", 10))
		Expect(answer()).To(BeNumerically("<=", 99))
	})
})
