package bdd

import (
	"fmt"
	"github.com/onsi/gomega"
)

type panicMatcher struct {
	*MatcherInfo
}

// Panics succeeds if actual is a function that, when invoked, panics.
// Actual must be a function that takes no arguments and returns no results.
var Panics MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.Panic().Match(actual)
	},
}

// HasOccurred succeeds if actual is a non-nil error
// The typical Go error checking pattern looks like:
//
//  err := SomethingThatMightFail()
//  Check(err, HasOccurred)
var HasOccurred MatcherFactory = &MatcherInfo{
	Matcher: func(actual interface{}, _ []interface{}) (success bool, message string, err error) {
		return gomega.HaveOccured().Match(actual)
	},
}

var ErrorContains MatcherFactory = &MatcherInfo{
	Parameters: 1,
	Matcher: func(actual interface{}, expected []interface{}) (success bool, message string, err error) {
		errVal, ok := actual.(error)
		if !ok {
			err = fmt.Errorf("did not obtain error")
			return
		}
		substr, ok := expected[0].(string)
		if !ok {
			err = fmt.Errorf("'substr' must be a string")
			return
		}
		args := expected[1:]
		return gomega.ContainSubstring(substr, args...).Match(errVal.Error())
	},
}

//type panicContainsChecker struct {
//	*CheckerInfo
//}
//
//var PanicContains Checker = &panicContainsChecker{
//	&CheckerInfo{Name: "PanicContains", Params: []string{"function", "expected"}},
//}
//
//func (checker *panicContainsChecker) Check(params []interface{}, names []string) (result bool, errmsg string) {
//	f := reflect.ValueOf(params[0])
//	if f.Kind() != reflect.Func || f.Type().NumIn() != 0 {
//		return false, "Function must take zero arguments"
//	}
//	defer func() {
//		// If the function has not panicked, then don't do the check.
//		if errmsg != "" {
//			return
//		}
//		obtained := recover()
//		names[0] = "panic"
//		if e, ok := obtained.(error); ok {
//			result = strings.Contains(e.Error(), params[1].(string))
//		} else if e, ok := obtained.(string); ok {
//			result = strings.Contains(e, params[1].(string))
//			return
//		} else {
//			errmsg = "Panic value is not a string or an error"
//			return
//		}
//	}()
//	f.Call(nil)
//	return false, "Function has not panicked"
//}
