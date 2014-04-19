package bdd

import (
	"github.com/onsi/gomega"
)

// IsAssignableTo succeeds if actual is assignable to the type of expected.
// It will return an error when one of the values is nil.
var IsAssignableTo Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsAssignableTo",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.BeAssignableToTypeOf(expected[0]), actual)
	},
}
