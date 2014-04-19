package bdd

import "testing"

var testmap = map[string]int{"A": 1, "B": 2}

var mapMatcherTests = []matcherTest{
	// HasKey
	{
		result(testmap, HasKey, "A"),
		Result{Success: true},
	},
	{
		result(testmap, HasKey, 1),
		Result{
			FailureMessage:        "Expected <map[string]int | len:2>: {\"A\": 1, \"B\": 2} to have key <int>: 1",
			NegatedFailureMessage: "Expected <map[string]int | len:2>: {\"A\": 1, \"B\": 2} not to have key <int>: 1",
		},
	},

	// HasKeys
	{
		result(testmap, HasKeys, "A", "B"),
		Result{Success: true},
	},
	{
		result(testmap, HasKeys, "A", "C"),
		Result{
			FailureMessage:        "Expected <map[string]int | len:2>: {\"A\": 1, \"B\": 2} to have key <string>: C",
			NegatedFailureMessage: "Expected <map[string]int | len:2>: {\"A\": 1, \"B\": 2} not to have key <string>: C",
		},
	},
}

func Test_Map(t *testing.T) {
	testMatchers(t, mapMatcherTests)
}
