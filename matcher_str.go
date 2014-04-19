package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"strings"
)

// HasSubstr succeeds if actual is a string or stringer that contains the
// passed-in substring.
var HasSubstr Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "HasSubstr",
	apply: func(actual interface{}, expected []interface{}) Result {
		substr, ok := toString(expected[0])
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		return resultFromGomega(gomega.ContainSubstring(substr), actual)
	},
}

// Regexp succeeds if actual is a string or stringer that matches the
// passed-in regexp.
var MatchesRegexp Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "MatchesRegexp",
	apply: func(actual interface{}, expected []interface{}) Result {
		regex, ok := toString(expected[0])
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		str, ok := toString(actual)
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		return resultFromGomega(gomega.MatchRegexp(regex), str)
	},
}

// Regexp succeeds if actual is a string or stringer that matches the
// has the passed-in suffix.
var HasSuffix Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "HasSuffix",
	apply: func(actual interface{}, expected []interface{}) Result {
		str, ok := toString(actual)
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		suffix, ok := toString(expected[0])
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if strings.HasSuffix(str, suffix) {
			r.Success = true
		} else {
			r.FailureMessage = format.Message(actual, " to have suffix ", expected...)
			r.NegatedFailureMessage = format.Message(actual, " not to have suffix ", expected...)
		}
		return r
	},
}

// Regexp succeeds if actual is a string or stringer that matches the
// has the passed-in prefix.
var HasPrefix Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "HasSuffix",
	apply: func(actual interface{}, expected []interface{}) Result {
		str, ok := toString(actual)
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		prefix, ok := toString(expected[0])
		if !ok {
			err := fmt.Errorf("Expected a string or stringer, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if strings.HasPrefix(str, prefix) {
			r.Success = true
		} else {
			r.FailureMessage = format.Message(actual, " to have prefix ", expected...)
			r.NegatedFailureMessage = format.Message(actual, " not to have prefix ", expected...)
		}
		return r
	},
}
