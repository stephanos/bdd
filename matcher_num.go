package bdd

import (
	"github.com/onsi/gomega"
)

var IsGreaterThan Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsGreaterThan",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically(">", expected[0]), actual)
	},
}

var IsGreaterThanOrEqTo Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsGreaterThanOrEqTo",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically(">=", expected[0]), actual)
	},
}

var IsLessThan Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsLessThan",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("<", expected[0]), actual)
	},
}

var IsLessThanOrEqTo Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsLessThanOrEqTo",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("<=", expected[0]), actual)
	},
}

var EqualsNum Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "EqualsNum",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("==", expected[0]), actual)
	},
}

var IsRoughly Matcher = &matcher{
	minArgs: 2,
	maxArgs: 2,
	name:    "IsRoughly",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("~", expected[0], expected[1]), actual)
	},
}
