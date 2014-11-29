package bdd

import (
	"strings"
	"testing"

	"github.com/onsi/gomega/format"
)

var openChan = make(chan bool)
var closedChan = make(chan bool)

var chanMatcherTests = []matcherTest{
	// IsClosed
	{
		result(openChan, IsClosed),
		Result{
			FailureMessage:        closedMessageFor(openChan, "closed"),
			NegatedFailureMessage: closedMessageFor(openChan, "open"),
		},
	},
	{
		result(closedChan, IsClosed),
		Result{Success: true},
	},
}

func Test_Chan(t *testing.T) {
	close(closedChan)

	testMatchers(t, chanMatcherTests)
}

func closedMessageFor(c chan bool, s string) string {
	return strings.Replace(
		strings.Replace(format.Message(c, " to be "+s), "\n", "", -1),
		"   ", "", -1)
}
