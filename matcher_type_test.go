package bdd

import "testing"

var typeMatcherTests = []matcherTest{
	// IsAssignableTo
	{
		result("string", IsAssignableTo, "string"),
		Result{Success: true},
	},
	{
		result(1, IsAssignableTo, "string"),
		Result{
			FailureMessage:        "Expected <int>: 1 to be assignable to the type: string",
			NegatedFailureMessage: "Expected <int>: 1 not to be assignable to the type: string",
		},
	},
}

func Test_Type(t *testing.T) {
	testMatchers(t, typeMatcherTests)
}
