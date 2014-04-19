package bdd

import "testing"

var notMatcherTests = []matcherTest{
	{
		result(1, Not(Equals), 0),
		Result{Success: true},
	},
	{
		result(1, Not(Not(Equals)), 1),
		Result{Success: true},
	},
	{
		result(1, Not(Equals), 1),
		Result{
			FailureMessage:        "Expected <int>: 1 not to equal <int>: 1",
			NegatedFailureMessage: "Expected <int>: 1 to equal <int>: 1",
		},
	},
}

func Test_Not(t *testing.T) {
	testMatchers(t, notMatcherTests)
}
