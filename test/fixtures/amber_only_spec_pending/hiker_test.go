package hiker

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// PIt marks the spec pending, so ginkgo registers it and runs nothing. The
// suite proves as little as one holding no specs at all, so the honest colour
// is amber.
var _ = Describe("answer", func() {

	PIt("is life the universe and everything", func() {
		Expect(answer()).To(Equal(42))
	})
})
