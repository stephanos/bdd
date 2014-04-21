package bdd

import (
	"fmt"
	"github.com/onsi/gomega/format"
	"time"
)

// IsSameTimeAs succeeds if actual is the same time or later
// than the passed-in time.
var IsSameTimeAs = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsSameTimeAs",
	apply: func(actual interface{}, expected []interface{}) Result {
		t1, ok := actual.(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		t2, ok := expected[0].(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if t1.Equal(t2) {
			r.Success = true
		} else {
			r.FailureMessage = timeMismatch(t1, " to be same time as ", t2)
			r.NegatedFailureMessage = timeMismatch(t1, " not to be same time as ", t2)
		}
		return r
	},
}

// IsBefore succeeds if actual is earlier than the passed-in time.
var IsBefore = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsBefore",
	apply: func(actual interface{}, expected []interface{}) Result {
		before, ok := actual.(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		after, ok := expected[0].(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if before.Before(after) {
			r.Success = true
		} else {
			r.FailureMessage = timeMismatch(before, " to be before ", after)
			r.NegatedFailureMessage = timeMismatch(before, " not to be before ", after)
		}
		return r
	},
}

// IsAfter succeeds if actual is later than the passed-in time.
var IsAfter = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsAfter",
	apply: func(actual interface{}, expected []interface{}) Result {
		after, ok := actual.(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		before, ok := expected[0].(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if after.After(before) {
			r.Success = true
		} else {
			r.FailureMessage = timeMismatch(after, " to be after ", before)
			r.NegatedFailureMessage = timeMismatch(after, " not to be after ", before)
		}
		return r
	},
}

// IsOnOrBefore succeeds if actual is the same time or earlier
// than the passed-in time.
var IsOnOrBefore = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsOnOrBefore",
	apply: func(actual interface{}, expected []interface{}) Result {
		before, ok := actual.(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		after, ok := expected[0].(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if before.Before(after) || before.Equal(after) {
			r.Success = true
		} else {
			r.FailureMessage = timeMismatch(before, " to be before or same time as ", after)
			r.NegatedFailureMessage = timeMismatch(before, " to be after ", after)
		}
		return r
	},
}

// IsOnOrAfter succeeds if actual is the same time or later
// than the passed-in time.
var IsOnOrAfter = &matcher{
	minArgs: 1,
	maxArgs: 1,
	name:    "IsOnOrAfter",
	apply: func(actual interface{}, expected []interface{}) Result {
		after, ok := actual.(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(actual, 1))
			return Result{Error: err}
		}

		before, ok := expected[0].(time.Time)
		if !ok {
			err := fmt.Errorf("expected a time.Time, got: \n %s", format.Object(expected[0], 1))
			return Result{Error: err}
		}

		var r Result
		if after.After(before) || after.Equal(before) {
			r.Success = true
		} else {
			r.FailureMessage = timeMismatch(after, " to be after or same time as ", before)
			r.NegatedFailureMessage = timeMismatch(after, " to be before ", before)
		}
		return r
	},
}

func timeMismatch(t1 time.Time, message string, t2 time.Time) string {
	strT1, strT2 := formatTime(t1), formatTime(t2)
	return fmt.Sprintf("Expected\n%s\n%s\n%s", strT1, message, strT2)
}

func formatTime(t time.Time) string {
	return format.Indent + t.String()
}
