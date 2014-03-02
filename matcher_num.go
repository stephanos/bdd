package bdd

import (
	"github.com/onsi/gomega"
)

var IsGreaterThan MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeNumerically(">", expected[0]).Match(actual)
	},
}

var IsGreaterThanOrEqTo MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeNumerically(">=", expected[0]).Match(actual)
	},
}

var IsLessThan MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeNumerically("<", expected[0]).Match(actual)
	},
}

var IsLessThanOrEqTo MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeNumerically("<=", expected[0]).Match(actual)
	},
}

var IsNum MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeNumerically("==", expected[0]).Match(actual)
	},
}

var IsRoughly MatcherFactory = &MatcherInfo{
	Parameters: 2,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeNumerically("~", expected[0], expected[1]).Match(actual)
	},
}
