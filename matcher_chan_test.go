package bdd

import (
	"strings"
	"testing"

	"github.com/onsi/gomega/format"
)

var openChan = make(chan bool)
var closedChan = make(chan bool)

func Test_Chan(t *testing.T) {
	close(closedChan)

	testMatchers(t, []matcherTest{
		// IsClosed
		{
			result(closedChan, IsClosed),
			Result{Success: true},
		},
		{
			result(openChan, IsClosed),
			Result{
				FailureMessage:        closedMessageFor(openChan, "closed"),
				NegatedFailureMessage: closedMessageFor(openChan, "open"),
			},
		},
	})
}

func closedMessageFor(c chan bool, s string) string {
	return strings.Replace(
		strings.Replace(format.Message(c, " to be "+s), "\n", "", -1),
		"   ", "", -1)
}
