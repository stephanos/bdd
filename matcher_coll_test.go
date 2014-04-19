package bdd

import (
	"errors"
	"testing"
)

var testcoll = []string{"A", "B", "C"}

var collMatcherTests = []matcherTest{
	// HasLen
	{
		result(testcoll, HasLen, 3),
		Result{Success: true},
	},
	{
		result(testmap, HasLen, 2),
		Result{Success: true},
	},
	{
		result("abc", HasLen, 3),
		Result{Success: true},
	},
	{
		result(testcoll, HasLen, 1),
		Result{
			FailureMessage:        "Expected <[]string | len:3, cap:3>: [\"A\", \"B\", \"C\"] to have length 1",
			NegatedFailureMessage: "Expected <[]string | len:3, cap:3>: [\"A\", \"B\", \"C\"] not to have length 1",
		},
	},

	// IsEmpty
	{
		result([]string{}, IsEmpty),
		Result{Success: true},
	},
	{
		result("", IsEmpty),
		Result{Success: true},
	},
	{
		result(make(map[int]string, 0), IsEmpty),
		Result{Success: true},
	},
	{
		result(make(chan string, 0), IsEmpty),
		Result{Success: true},
	},
	{
		result(testcoll, IsEmpty),
		Result{
			FailureMessage:        "Expected <[]string | len:3, cap:3>: [\"A\", \"B\", \"C\"] to be empty",
			NegatedFailureMessage: "Expected <[]string | len:3, cap:3>: [\"A\", \"B\", \"C\"] not to be empty",
		},
	},

	// HasElem
	{
		result(testcoll, HasElem, "A"),
		Result{Success: true},
	},
	{
		result(testmap, HasElem, 1),
		Result{Success: true},
	},
	{
		result([]string{}, HasElem, "A"),
		Result{
			FailureMessage:        "Expected <[]string | len:0, cap:0>: [] to contain element matching <string>: A",
			NegatedFailureMessage: "Expected <[]string | len:0, cap:0>: [] not to contain element matching <string>: A",
		},
	},

	// HasElems
	{
		result(testcoll, HasElems, "A", "B"),
		Result{Success: true},
	},
	{
		result(testmap, HasElems, 1, 2),
		Result{Success: true},
	},
	{
		result(testcoll, HasElems, "A", "D"),
		Result{
			FailureMessage:        "Expected <[]string | len:3, cap:3>: [\"A\", \"B\", \"C\"] to contain element matching <string>: D",
			NegatedFailureMessage: "Expected <[]string | len:3, cap:3>: [\"A\", \"B\", \"C\"] not to contain element matching <string>: D",
		},
	},

	// Contains
	{
		result("ABC", Contains, "A", "B"),
		Result{Success: true},
	},
	{
		result(testcoll, Contains, "A", "B"),
		Result{Success: true},
	},
	{
		result(testmap, Contains, 1, 2),
		Result{Success: true},
	},
	{
		result(errors.New("database error"), Contains, "database", "error"),
		Result{Success: true},
	},
}

func Test_Coll(t *testing.T) {
	testMatchers(t, collMatcherTests)
}
