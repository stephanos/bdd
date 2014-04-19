package bdd

import (
	"fmt"
	"strings"
	"testing"
)

type matcherTest struct {
	actual   testResult
	expected Result
}

func testMatchers(t *testing.T, matcherTests []matcherTest) {
	for _, test := range matcherTests {
		actual := test.actual
		expected := test.expected

		if actual.Success != expected.Success {
			if actual.Success {
				t.Errorf("expected %s to be unsuccessful", actual)
			} else {
				t.Errorf("expected %s to be successful", actual)
			}
		}

		actualErr := flatten(fmt.Sprintf("%v", actual.Error))
		expectedErr := flatten(fmt.Sprintf("%v", expected.Error))
		if actualErr != expectedErr {
			if actual.Error == nil {
				t.Errorf("expected %s to run without error, but got: \n   %q", actual, actualErr)
			} else {
				t.Errorf("expected %s to return error \n %q, \n    but got \n %q", actual, expectedErr, actualErr)
			}
		}

		if !expected.Success {
			actualFailureMessage := flatten(actual.FailureMessage)
			if actualFailureMessage != expected.FailureMessage {
				t.Errorf("expected %s to return failure message \n %q, \n    but got \n %q",
					actual, expected.FailureMessage, actualFailureMessage)
			}

			actualNegatedFailureMessage := flatten(actual.NegatedFailureMessage)
			if actualNegatedFailureMessage != expected.NegatedFailureMessage {
				t.Errorf("expected %s to return negated failure message \n %q, \n    but got \n %q",
					actual, expected.NegatedFailureMessage, actualNegatedFailureMessage)
			}
		}
	}
}

// ==== HELPERS

type testResult struct {
	Result

	obtained interface{}
	matcher  Matcher
	args     []interface{}
}

func result(obtained interface{}, matcher Matcher, args ...interface{}) testResult {
	return testResult{newChecker(obtained).result(matcher, args), obtained, matcher, args}
}

func flatten(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func (r testResult) String() string {
	var matcherName = "?"
	if nm, ok := r.matcher.(namedMatcher); ok {
		matcherName = nm.Name()
	}

	lenArgs := len(r.args)
	if lenArgs == 0 {
		return fmt.Sprintf("[%v %s]", r.obtained, matcherName)
	} else if lenArgs == 1 {
		return fmt.Sprintf("[%v %s %v]", r.obtained, matcherName, r.args[0])
	} else {
		return fmt.Sprintf("[%v %s %v]", r.obtained, matcherName, r.args)
	}
}
