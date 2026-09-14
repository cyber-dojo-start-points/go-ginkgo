package hiker

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHiker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hiker Suite")
}
