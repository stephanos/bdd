package bdd

import (
	"errors"
	"testing"
)

var errMatcherTests = []matcherTest{
	// Panics
	{
		result(func() {
			panic("stay calm! not!")
		}, Panics),
		Result{Success: true},
	},

	// HasOccurred
	{
		result(errors.New("an error"), HasOccurred),
		Result{Success: true},
	},
	{
		result(nil, HasOccurred),
		Result{
			FailureMessage:        "expected an error to have occured. Got: <nil>: nil",
			NegatedFailureMessage: "expected an error to have occured. Got: <nil>: nil",
		},
	},

	// ErrorContains
	{
		result(errors.New("database error"), ErrorContains, "database"),
		Result{Success: true},
	},
	{
		result(nil, ErrorContains, "error"),
		Result{
			Error: errors.New("expected an error, got: <nil>: nil"),
		},
	},
	{
		result(errors.New("database error"), ErrorContains, 42),
		Result{
			Error: errors.New("expected a string, got: <int>: 42"),
		},
	},
	{
		result(errors.New("database error"), ErrorContains, "HTTP"),
		Result{
			FailureMessage:        "Expected <error>: database error to contain <string>: HTTP",
			NegatedFailureMessage: "Expected <error>: database error not to contain <string>: HTTP",
		},
	},
}

func Test_Err(t *testing.T) {
	testMatchers(t, errMatcherTests)
}
