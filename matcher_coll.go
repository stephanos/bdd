package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
)

// HasLen succeeds if actual has the passed-in length.
// Actual must be of type string, array, map, chan, or slice.
var HasLen MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		len, ok := expected[0].(int)
		if !ok {
			err = fmt.Errorf("'length' must be an int")
			return
		}
		return gomega.HaveLen(len).Match(actual)
	},
}

// IsEmpty succeeds if actual is empty. Actual must be of type string, array, map, chan, or slice.
var IsEmpty MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.BeEmpty().Match(actual)
	},
}

// HasElem succeeds if actual contains the passed in element.
// By default ContainsItem() uses Equal() to perform the match, however a
// matcher can be passed in instead:
//
//  Check([]string{"Foo", "FooBar"}, HasElem, ContainString("Bar"))
//
// Actual must be an array, slice or map.
// For maps, containElement searches through the map's values.
var HasElem MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.ContainElement(expected[0]).Match(actual)
	},
}

var Contains MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		switch actual := actual.(type) {
		case string, fmt.Stringer:
			return HasSubstr.New(expected).Match(actual)
		case error:
			return ErrorContains.New(expected).Match(actual)
		}

		return HasElem.New(expected).Match(actual)
	},
}
