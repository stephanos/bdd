package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
	"strings"
)

// Regexp succeeds if actual is a string or stringer that matches the
// passed-in regexp. Optional arguments can be provided to construct a regexp via fmt.Sprintf().
var MatchesRegexp MatcherFactory = &MatcherInfo{
	Parameters: -1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		regex, ok := expected[0].(string)
		if !ok {
			err = fmt.Errorf("'regexp' must be a string")
			return
		}

		args := expected[1:]
		return gomega.MatchRegexp(regex, args...).Match(actual)
	},
}

// Substr succeeds if actual is a string or stringer that contains the
// passed-in regexp. Optional arguments can be provided to construct the substring via fmt.Sprintf().
var HasSubstr MatcherFactory = &MatcherInfo{
	Parameters: -1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		substr, ok := expected[0].(string)
		if !ok {
			err = fmt.Errorf("'substr' must be a string")
			return
		}

		args := expected[1:]
		return gomega.ContainSubstring(substr, args...).Match(actual)
	},
}

var HasSuffix MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		suffix, ok := expected[0].(string)
		if !ok {
			err = fmt.Errorf("'suffix' must be a string")
			return
		}

		success, message, err = gomega.ContainSubstring(suffix).Match(actual)
		if success && !strings.HasSuffix(actual.(string), suffix) {
			success = false
		}
		return
	},
}

var HasPrefix MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		prefix, ok := expected[0].(string)
		if !ok {
			err = fmt.Errorf("'prefix' must be a string")
			return
		}

		success, message, err = gomega.ContainSubstring(prefix).Match(actual)
		if success && !strings.HasPrefix(actual.(string), prefix) {
			success = false
		}
		return
	},
}
