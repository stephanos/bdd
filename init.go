package bdd

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func RunSpecs(t *testing.T, descr string) {

	// Ginkgo test signals failure by calling 'Fail' function - passing this function to Gomega
	gomega.RegisterFailHandler(func(message string, callerSkip ...int) {
		skip := 3
		if len(callerSkip) > 0 {
			skip += callerSkip[0]
		}
		ginkgo.Fail(message, skip)
	})

	// tell Ginkgo to start the test suite
	ginkgo.RunSpecs(t, descr)
}
