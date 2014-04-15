package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
)

type panicMatcher struct {
	*MatcherInfo
}

// Panics succeeds if actual is a function that, when invoked, panics.
// Actual must be a function that takes no arguments and returns no results.
var Panics MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.Panic().Match(actual)
	},
}

// HasOccurred succeeds if actual is a non-nil error
// The typical Go error checking pattern looks like:
//
//  err := SomethingThatMightFail()
//  Check(err, HasOccurred)
var HasOccurred MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.HaveOccurred().Match(actual)
	},
}

var ErrorContains MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		errVal, ok := actual.(error)
		if !ok {
			err = fmt.Errorf("did not obtain error")
			return
		}
		substr, ok := expected[0].(string)
		if !ok {
			err = fmt.Errorf("'substr' must be a string")
			return
		}
		args := expected[1:]
		return gomega.ContainSubstring(substr, args...).Match(errVal.Error())
	},
}
