package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"strings"
)

type panicMatcher struct {
	*matcher
}

// Panics succeeds if actual is a function that, when invoked, panics.
// Actual must be a function that takes no arguments and returns no results.
var Panics Matcher = &matcher{
	name: "Panics",
	apply: func(actual interface{}, _ []interface{}) Result {
		return resultFromGomega(gomega.Panic(), actual)
	},
}

// HasOccurred succeeds if actual is a non-nil error.
var HasOccurred Matcher = &matcher{
	name: "HasOccurred",
	apply: func(actual interface{}, _ []interface{}) Result {
		if actual == nil {
			return Result{
				FailureMessage:        fmt.Sprintf("Expected an error to have occured.  Got:\n%s", format.Object(actual, 1)),
				NegatedFailureMessage: fmt.Sprintf("Expected an error to have occured. Got:\n%s", format.Object(actual, 1)),
			}
		}
		return resultFromGomega(gomega.HaveOccurred(), actual)
	},
}

// ErrorContains succeeds if actual is a non-nil error and contains
// the passed-in substring.
var ErrorContains Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "ErrorContains",
	apply: func(actual interface{}, expected []interface{}) Result {
		err, ok := actual.(error)
		if !ok {
			err := fmt.Errorf("Expected an error, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}
		errStr := err.Error()

		substr, ok := toString(expected[0])
		if !ok {
			err := fmt.Errorf("Expected a string, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if strings.Contains(errStr, substr) {
			r.Success = true
		} else {
			r.FailureMessage = fmt.Sprintf("Expected\n<error>: %s\n%s\n%s", errStr, " to contain ", format.Object(expected[0], 1))
			r.NegatedFailureMessage = fmt.Sprintf("Expected\n<error>: %s\n%s\n%s", errStr, " not to contain ", format.Object(expected[0], 1))
		}
		return r
	},
}
