package bdd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

var globalFailHandler types.GomegaFailHandler

// RunSpecs is the entry point for the test runner.
// You must call this within a Golang testing TestX(t *testing.T) function.
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
