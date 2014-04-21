package bdd

import (
	"errors"
	"testing"
	"time"
)

var (
	moonLanding    = time.Unix(-14183900, 0)
	fallBerlinWall = time.Unix(626644800, 0)
)

var timeMatcherTests = []matcherTest{
	// IsBefore
	{
		result(moonLanding, IsBefore, fallBerlinWall),
		Result{Success: true},
	},
	{
		result(fallBerlinWall, IsBefore, 1963),
		Result{
			Error: errors.New("expected a time.Time, got: <int>: 1963"),
		},
	},
	{
		result(1989, IsBefore, moonLanding),
		Result{
			Error: errors.New("expected a time.Time, got: <int>: 1989"),
		},
	},
	{
		result(fallBerlinWall, IsBefore, moonLanding),
		Result{
			FailureMessage:        "Expected 1989-11-09 21:00:00 +0100 CET to be before 1969-07-20 21:01:40 +0100 CET",
			NegatedFailureMessage: "Expected 1989-11-09 21:00:00 +0100 CET not to be before 1969-07-20 21:01:40 +0100 CET",
		},
	},

	// IsAfter
	{
		result(fallBerlinWall, IsAfter, moonLanding),
		Result{Success: true},
	},
	{
		result(moonLanding, IsAfter, fallBerlinWall),
		Result{
			FailureMessage:        "Expected 1969-07-20 21:01:40 +0100 CET to be after 1989-11-09 21:00:00 +0100 CET",
			NegatedFailureMessage: "Expected 1969-07-20 21:01:40 +0100 CET not to be after 1989-11-09 21:00:00 +0100 CET",
		},
	},

	// IsOnOrBefore
	{
		result(moonLanding, IsOnOrBefore, fallBerlinWall),
		Result{Success: true},
	},
	{
		result(moonLanding, IsOnOrBefore, moonLanding),
		Result{Success: true},
	},
	{
		result(fallBerlinWall, IsOnOrBefore, moonLanding),
		Result{
			FailureMessage:        "Expected 1989-11-09 21:00:00 +0100 CET to be before or same time as 1969-07-20 21:01:40 +0100 CET",
			NegatedFailureMessage: "Expected 1989-11-09 21:00:00 +0100 CET to be after 1969-07-20 21:01:40 +0100 CET",
		},
	},

	// IsOnOrAfter
	{
		result(fallBerlinWall, IsOnOrAfter, moonLanding),
		Result{Success: true},
	},
	{
		result(fallBerlinWall, IsOnOrAfter, fallBerlinWall),
		Result{Success: true},
	},
	{
		result(moonLanding, IsOnOrAfter, fallBerlinWall),
		Result{
			FailureMessage:        "Expected 1969-07-20 21:01:40 +0100 CET to be after or same time as 1989-11-09 21:00:00 +0100 CET",
			NegatedFailureMessage: "Expected 1969-07-20 21:01:40 +0100 CET to be before 1989-11-09 21:00:00 +0100 CET",
		},
	},
}

func Test_Time(t *testing.T) {
	testMatchers(t, timeMatcherTests)
}
