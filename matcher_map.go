package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
)

// Key succeeds if actual is a map with the passed in key.
// By default Key uses Equal() to perform the match, however a matcher can be passed in instead.
var HasKey MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		return gomega.HaveKey(expected[0]).Match(actual)
	},
}

var HasKeys MatcherFactory = &MatcherInfo{
	Parameters: -1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		for _, exp := range expected {
			key, ok := exp.(string)
			if !ok {
				err = fmt.Errorf("'key' must be string")
				return
			}

			success, message, err = gomega.HaveKey(key).Match(actual)
			if !success {
				return
			}
		}
		return
	},
}
