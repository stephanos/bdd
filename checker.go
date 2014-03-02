package bdd

import (
	"github.com/onsi/gomega"
)

type Checker struct {
	gomega.Actual

	failed   bool
	negated  bool
	obtained interface{}
}

func newChecker(obtained interface{}) *Checker {
	return &Checker{
		obtained: obtained,
		Actual:   gomega.Expect(obtained),
	}
}

// Ω wraps an actual value allowing assertions to be made on it:
//        Ω("foo", Equals, "foo")
func Ω(obtained interface{}, matcher MatcherFactory, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

// Assert wraps an actual value allowing assertions to be made on it:
//        Expect("foo", Equals, "foo")
func Assert(obtained interface{}, matcher MatcherFactory, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

// Check wraps an actual value allowing assertions to be made on it:
//        Check("foo", Equals, "foo")
func Check(obtained interface{}, matcher MatcherFactory, args ...interface{}) *Checker {
	return newChecker(obtained).run(matcher, args...)
}

func (chk *Checker) And(matcher MatcherFactory, args ...interface{}) *Checker {
	chk.negated = false
	if !chk.failed {
		return chk.run(matcher, args...)
	}
	return chk
}

func (chk *Checker) run(matcher MatcherFactory, args ...interface{}) *Checker {
	if notMatcher, ok := matcher.(*NotMatcher); ok {
		chk.negated = !chk.negated
		matcher = notMatcher.inner
	}

	inst := matcher.New(args)
	if chk.negated {
		chk.failed = !chk.Actual.ToNot(inst)
	} else {
		chk.failed = !chk.Actual.To(inst)
	}
	return chk
}
