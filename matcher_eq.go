package bdd

import (
	"github.com/onsi/gomega"
)

// Equal uses reflect.DeepEqual to compare actual with expected.
// It's strict about types when performing comparisons.
// It is an error for both actual and expected to be nil. Use IsNil instead.
var Equals Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "Equals",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.Equal(expected[0]), actual)
	},
}

// IsEquivalentTo is more lax than Equal, allowing equality between different types.
// This is done by converting actual to have the type of expected before
// attempting equality with reflect.DeepEqual.
// It is an error for actual and expected to be nil. Use IsNil instead.
var IsEquivalentTo Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsEquivalentTo",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeEquivalentTo(expected[0]), actual)
	},
}

// IsNil succeeds if actual is nil
var IsNil Matcher = &matcher{

	name: "IsNil",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeNil(), actual)
	},
}

// NotNil succeeds if actual is not nil
var NotNil Matcher = Not(IsNil)

// IsTrue succeeds if actual is true
var IsTrue Matcher = &matcher{

	name: "IsTrue",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeTrue(), actual)
	},
}

// IsFalse succeeds if actual is false
var IsFalse Matcher = &matcher{

	name: "IsFalse",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeFalse(), actual)
	},
}

// IsZero succeeds if actual is the zero value for its type or if actual is nil.
var IsZero Matcher = &matcher{

	name: "IsZero",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeZero(), actual)
	},
}
