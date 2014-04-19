package bdd

import (
	"errors"
	"testing"
)

var eqMatcherTests = []matcherTest{
	// Equals
	{
		result(1, Equals, 1),
		Result{Success: true},
	},
	{
		result(1, Equals),
		Result{
			Error: errors.New("Expected at least 1 parameter(s), but got 0"),
		},
	},
	{
		result(1, Equals, 1, 2),
		Result{
			Error: errors.New("Expected at most 1 parameter(s), but got 2"),
		},
	},
	{
		result(1, Equals, 0),
		Result{
			FailureMessage:        "Expected <int>: 1 to equal <int>: 0",
			NegatedFailureMessage: "Expected <int>: 1 not to equal <int>: 0",
		},
	},

	// IsEquivalentTo
	{
		result(1.0, IsEquivalentTo, 1),
		Result{Success: true},
	},
	{
		result(2.0, IsEquivalentTo, 1),
		Result{
			FailureMessage:        "Expected <float64>: 2 to be equivalent to <int>: 1",
			NegatedFailureMessage: "Expected <float64>: 2 not to be equivalent to <int>: 1",
		},
	},

	// IsNil
	{
		result(nil, IsNil),
		Result{Success: true},
	},
	{
		result(1, IsNil),
		Result{
			FailureMessage:        "Expected <int>: 1 to be nil",
			NegatedFailureMessage: "Expected <int>: 1 not to be nil",
		},
	},
	{
		result(nil, IsNil, nil),
		Result{
			Error: errors.New("Expected at most 0 parameter(s), but got 1"),
		},
	},

	// NotNil
	{
		result(1, NotNil),
		Result{Success: true},
	},
	{
		result(nil, NotNil),
		Result{
			FailureMessage:        "Expected <nil>: nil not to be nil",
			NegatedFailureMessage: "Expected <nil>: nil to be nil",
		},
	},

	// IsFalse
	{
		result(false, IsFalse),
		Result{Success: true},
	},
	{
		result(true, IsFalse),
		Result{
			FailureMessage:        "Expected <bool>: true to be false",
			NegatedFailureMessage: "Expected <bool>: true not to be false",
		},
	},

	// IsTrue
	{
		result(true, IsTrue),
		Result{Success: true},
	},
	{
		result(false, IsTrue),
		Result{
			FailureMessage:        "Expected <bool>: false to be true",
			NegatedFailureMessage: "Expected <bool>: false not to be true",
		},
	},

	// IsZero
	{
		result("", IsZero),
		Result{Success: true},
	},
	{
		result("1", IsZero),
		Result{
			FailureMessage:        "Expected <string>: 1 to be zero-valued",
			NegatedFailureMessage: "Expected <string>: 1 not to be zero-valued",
		},
	},
}

func Test_Equals(t *testing.T) {
	testMatchers(t, eqMatcherTests)
}
