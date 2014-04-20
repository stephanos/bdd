package bdd

import (
	"errors"
	"testing"
)

var strMatcherTests = []matcherTest{
	// HasSubstr
	{
		result("golang", HasSubstr, "ola"),
		Result{Success: true},
	},
	{
		result("team", HasSubstr, 1),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 1"),
		},
	},
	{
		result("team", HasSubstr, "I"),
		Result{
			FailureMessage:        "Expected <string>: team to contain substring <string>: I",
			NegatedFailureMessage: "Expected <string>: team not to contain substring <string>: I",
		},
	},

	// HasSuffix
	{
		result("golang", HasSuffix, "lang"),
		Result{Success: true},
	},
	{
		result("golang", HasSuffix, 1),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 1"),
		},
	},
	{
		result(11, HasSuffix, "one"),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 11"),
		},
	},
	{
		result("golang", HasSuffix, "go"),
		Result{
			FailureMessage:        "Expected <string>: golang to have suffix <string>: go",
			NegatedFailureMessage: "Expected <string>: golang not to have suffix <string>: go",
		},
	},

	// HasPrefix
	{
		result("golang", HasPrefix, "go"),
		Result{Success: true},
	},
	{
		result("golang", HasPrefix, 1),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 1"),
		},
	},
	{
		result(11, HasPrefix, "ten"),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 11"),
		},
	},
	{
		result("golang", HasPrefix, "lang"),
		Result{
			FailureMessage:        "Expected <string>: golang to have prefix <string>: lang",
			NegatedFailureMessage: "Expected <string>: golang not to have prefix <string>: lang",
		},
	},

	// MatchesRegexp
	{
		result("golang", MatchesRegexp, "go*"),
		Result{Success: true},
	},
	{
		result("golang", MatchesRegexp, 1),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 1"),
		},
	},
	{
		result(1, MatchesRegexp, "golang"),
		Result{
			Error: errors.New("expected a string or stringer, got: <int>: 1"),
		},
	},
	{
		result("golang", MatchesRegexp, "[A-Z]+"),
		Result{
			FailureMessage:        "Expected <string>: golang to match regular expression <string>: [A-Z]+",
			NegatedFailureMessage: "Expected <string>: golang not to match regular expression <string>: [A-Z]+",
		},
	},
}

func Test_Str(t *testing.T) {
	testMatchers(t, strMatcherTests)
}
