package bdd

import (
	"github.com/onsi/gomega"
)

// HasKey succeeds if actual is a map with the passed-in key.
var HasKey = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "HasKey",
	apply: func(actual interface{}, expected []interface{}) Result {
		return resultFromGomega(gomega.HaveKey(expected[0]), actual)
	},
}

// HasKeys succeeds if actual is a map with all passed-in keys.
var HasKeys = &matcher{
	minArgs: 1,
	maxArgs: 1<<(31) - 1,
	name:    "HasKeys",
	apply: func(actual interface{}, expected []interface{}) (r Result) {
		for _, key := range expected {
			r = resultFromGomega(gomega.HaveKey(key), actual)
			if !r.Success {
				return
			}
		}
		return
	},
}
