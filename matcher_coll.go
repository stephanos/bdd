package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

// HasLen succeeds if actual has the passed-in length. Actual must be of type
// string, array, map, chan, or slice.
var HasLen Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "HasLen",
	apply: func(actual interface{}, expected []interface{}) Result {
		len, ok := expected[0].(int)
		if !ok {
			err := fmt.Errorf("Expected length to be an int, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		return resultFromGomega(gomega.HaveLen(len), actual)
	},
}

// IsEmpty succeeds if actual is empty. Actual must be of type
// string, array, map, chan, or slice.
var IsEmpty Matcher = &matcher{
	name: "IsEmpty",
	apply: func(actual interface{}, _ []interface{}) Result {
		return resultFromGomega(gomega.BeEmpty(), actual)
	},
}

// HasElem succeeds if actual contains the passed in element. Actual must be
// an array, slice or map. For maps, HasElem searches through the map's values.
var HasElem Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "HasElem",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.ContainElement(expected[0]), actual)
	},
}

// HasElems succeeds if actual contains all passed in element. Actual must be
// an array, slice or map. For maps, HasElems searches through the map's values.
var HasElems Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1<<(31) - 1,
	name:    "HasElems",
	apply: func(actual interface{}, expected []interface{}) (r Result) {
		for _, val := range expected {
			r = resultFromGomega(gomega.ContainElement(val), actual)
			if !r.Success {
				return
			}
		}
		return
	},
}

// Contains succeeds if actual contains all passed in substrings / elements.
// Actual must be an error, string, stringer, array, slice or map.
var Contains Matcher = &matcher{
	minArgs: 1,
	maxArgs: 1<<(31) - 1,
	name:    "Contains",
	apply: func(actual interface{}, expected []interface{}) (r Result) {

		var apply func(obtained interface{}, args []interface{}) Result
		switch actual.(type) {
		case string, fmt.Stringer:
			apply = HasSubstr.Apply
		case error:
			apply = ErrorContains.Apply
		default:
			apply = HasElem.Apply
		}

		for _, exp := range expected {
			r = apply(actual, []interface{}{exp})
			if !r.Success {
				return
			}
		}
		return
	},
}
