package bdd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

var globalFailHandler gomega.OmegaFailHandler

func RunSpecs(t *testing.T, descr string) {

	globalFailHandler = func(message string, callerSkip ...int) {
		skip := 3
		if len(callerSkip) > 0 {
			skip += callerSkip[0]
		}
		ginkgo.Fail(message, skip)
	}

	// Ginkgo test signals failure by calling 'Fail' function - passing this function to Gomega
	gomega.RegisterFailHandler(globalFailHandler)

	// tell Ginkgo to start the test suite
	ginkgo.RunSpecs(t, descr)
}
