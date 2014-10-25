package bdd

import "testing"

var testmap1 = map[string]int{"A": 1}
var testmap2 = map[string]int{"A": 1, "B": 2}

var mapMatcherTests = []matcherTest{
	// HasKey
	{
		result(testmap2, HasKey, "A"),
		Result{Success: true},
	},
	{
		result(testmap1, HasKey, "C"),
		Result{
			FailureMessage:        "Expected <map[string]int | len:1>: {\"A\": 1} to have key <string>: C",
			NegatedFailureMessage: "Expected <map[string]int | len:1>: {\"A\": 1} not to have key <string>: C",
		},
	},

	// HasKeys
	{
		result(testmap2, HasKeys, "A", "B"),
		Result{Success: true},
	},
	{
		result(testmap1, HasKeys, "A", "C"),
		Result{
			FailureMessage:        "Expected <map[string]int | len:1>: {\"A\": 1} to have key <string>: C",
			NegatedFailureMessage: "Expected <map[string]int | len:1>: {\"A\": 1} not to have key <string>: C",
		},
	},
}

func Test_Map(t *testing.T) {
	testMatchers(t, mapMatcherTests)
}
