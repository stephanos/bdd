package bdd

import (
	"fmt"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Checker makes assertions based on values and matchers.
type Checker struct {
	failed      bool
	lastMatcher Matcher
	obtained    interface{}
}

func newChecker(obtained interface{}) *Checker {
	return &Checker{
		obtained: obtained,
	}
}

// Ω wraps an actual value allowing assertions to be made on it:
//        Ω("foo", Equals, "foo")
func Ω(obtained interface{}, matcher Matcher, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

// Assert wraps an actual value allowing assertions to be made on it:
//        Assert("foo", Equals, "foo")
func Assert(obtained interface{}, matcher Matcher, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

// Expect wraps an actual value allowing assertions to be made on it:
//        Expect("foo", Equals, "foo")
func Expect(obtained interface{}, matcher Matcher, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

// Check wraps an actual value allowing assertions to be made on it:
//        Check("foo", Equals, "foo")
func Check(obtained interface{}, matcher Matcher, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

// CheckFail will record a failure for the current space and panic. This stops
// the current spec in its tracks - no subsequent assertions will be called.
func CheckFail(msg string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(msg, args...))
}

// And runs an assertion after the previous one. The obtained value is taken from
// the previous one. It is only run when the previous assertion was successful.
func (chk *Checker) And(args ...interface{}) *Checker {
	if !chk.failed {
		if len(args) < 1 {
			panic("missing arguments for And(...)")
		}

		if matcher, ok := args[0].(Matcher); ok {
			return chk.run(matcher, args[1:]...)
		}

		return chk.run(chk.lastMatcher, args...)
	}
	return chk
}

// ButNot runs an assertion after a previous one. The obtained value is taken from
// the previous one, the previous matcher negated. It is only run when the previous
// assertion was successful.
func (chk *Checker) ButNot(args ...interface{}) *Checker {
	if !chk.failed {
		return chk.run(Not(chk.lastMatcher), args...)
	}
	return chk
}

func (chk *Checker) result(matcher Matcher, args []interface{}) Result {
	chk.lastMatcher = matcher

	haveArgs := len(args)
	if m, ok := matcher.(MatcherArgsLimiter); ok {
		if haveArgs > m.MaxArgs() {
			err := fmt.Errorf("expected at most %d parameter(s), but got %d", m.MaxArgs(), haveArgs)
			return Result{Error: err}
		}
	}
	if m, ok := matcher.(MatcherArgsRequierer); ok {
		if haveArgs < m.MinArgs() {
			err := fmt.Errorf("expected at least %d parameter(s), but got %d", m.MinArgs(), haveArgs)
			return Result{Error: err}
		}
	}

	return matcher.Apply(chk.obtained, args)
}

func (chk *Checker) run(matcher Matcher, args ...interface{}) *Checker {
	gomegaMatcher := newGomegaMatcher(chk.result(matcher, args))
	chk.failed = !gomega.Expect(chk.obtained).To(gomegaMatcher)
	return chk
}
