package bdd

import (
	"github.com/onsi/gomega"
)

// IsGreaterThan succeeds if actual is a greater than the passed-in number.
var IsGreaterThan = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsGreaterThan",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically(">", expected[0]), actual)
	},
}

// IsGreaterThanOrEqTo succeeds if actual is a greater than or equal to
// the passed-in number.
var IsGreaterThanOrEqTo = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsGreaterThanOrEqTo",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically(">=", expected[0]), actual)
	},
}

// IsLessThan succeeds if actual is a smaller than the passed-in number.
var IsLessThan = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsLessThan",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("<", expected[0]), actual)
	},
}

// IsLessThanOrEqTo succeeds if actual is a smaller than or equal to
// the passed-in number.
var IsLessThanOrEqTo = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsLessThanOrEqTo",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("<=", expected[0]), actual)
	},
}

// EqualsNum succeeds if actual has the same numeric value as the passed-in number.
var EqualsNum = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "EqualsNum",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("==", expected[0]), actual)
	},
}

// IsRoughly succeeds if actual has about the same numeric value as the
// passed-in number. The second passed-in argument defines the threshold.
var IsRoughly = &matcher{
	minArgs: 2,
	maxArgs: 2,
	name:    "IsRoughly",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNumerically("~", expected[0], expected[1]), actual)
	},
}
