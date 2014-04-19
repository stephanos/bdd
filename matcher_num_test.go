package bdd

import "testing"

var numMatcherTests = []matcherTest{
	// IsGreaterThan
	{
		result(10, IsGreaterThan, 5),
		Result{Success: true},
	},
	{
		result(5, IsGreaterThan, 10),
		Result{
			FailureMessage:        "Expected <int>: 5 to be > <int>: 10",
			NegatedFailureMessage: "Expected <int>: 5 not to be > <int>: 10",
		},
	},

	// IsLessThan
	{
		result(5, IsLessThan, 10),
		Result{Success: true},
	},
	{
		result(10, IsLessThan, 5),
		Result{
			FailureMessage:        "Expected <int>: 10 to be < <int>: 5",
			NegatedFailureMessage: "Expected <int>: 10 not to be < <int>: 5",
		},
	},

	// IsGreaterThanOrEqTo
	{
		result(5, IsGreaterThanOrEqTo, 5),
		Result{Success: true},
	},
	{
		result(10, IsGreaterThanOrEqTo, 5),
		Result{Success: true},
	},
	{
		result(5, IsGreaterThanOrEqTo, 10),
		Result{
			FailureMessage:        "Expected <int>: 5 to be >= <int>: 10",
			NegatedFailureMessage: "Expected <int>: 5 not to be >= <int>: 10",
		},
	},

	// IsLessThanOrEqTo
	{
		result(5, IsLessThanOrEqTo, 5),
		Result{Success: true},
	},
	{
		result(5, IsLessThanOrEqTo, 10),
		Result{Success: true},
	},
	{
		result(10, IsLessThanOrEqTo, 5),
		Result{
			FailureMessage:        "Expected <int>: 10 to be <= <int>: 5",
			NegatedFailureMessage: "Expected <int>: 10 not to be <= <int>: 5",
		},
	},

	// EqualsNum
	{
		result(4, EqualsNum, 4.0),
		Result{Success: true},
	},
	{
		result(5, EqualsNum, 10),
		Result{
			FailureMessage:        "Expected <int>: 5 to be == <int>: 10",
			NegatedFailureMessage: "Expected <int>: 5 not to be == <int>: 10",
		},
	},

	// IsRoughly
	{
		result(10, IsRoughly, 9, 1),
		Result{Success: true},
	},
	{
		result(10, IsRoughly, 9, 0.5),
		Result{
			FailureMessage:        "Expected <int>: 10 to be ~ <int>: 9",
			NegatedFailureMessage: "Expected <int>: 10 not to be ~ <int>: 9",
		},
	},
}

func Test_Num(t *testing.T) {
	testMatchers(t, numMatcherTests)
}
