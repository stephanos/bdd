package bdd

import (
	"github.com/onsi/gomega"
)

type equalsMatcher struct {
	*MatcherInfo
}

// Equal uses reflect.DeepEqual to compare actual with expected.
// It's strict about types when performing comparisons.
// It is an error for both actual and expected to be nil. Use Nil instead.
var Equals MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.Equal(expected[0]).Match(actual)
	},
}

type equivalentMatcher struct {
	*MatcherInfo
}

// IsEquivalentTo is more lax than Equal, allowing equality between different types.
// This is done by converting actual to have the type of expected before
// attempting equality with reflect.DeepEqual.
// It is an error for actual and expected to be nil. Use BeNil() instead.
var IsEquivalentTo MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeEquivalentTo(expected[0]).Match(actual)
	},
}

// IsNil succeeds if actual is nil
var IsNil MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.BeNil().Match(actual)
	},
}

// NotNil succeeds if actual is not nil
var NotNil MatcherFactory = Not(IsNil)

// IsTrue succeeds if actual is true
var IsTrue MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.BeTrue().Match(actual)
	},
}

// IsFalse succeeds if actual is false
var IsFalse MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.BeFalse().Match(actual)
	},
}

// IsZero succeeds if actual is the zero value for its type or if actual is nil.
var IsZero MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.BeZero().Match(actual)
	},
}
