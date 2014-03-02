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

// FACTORY ========================================================================================

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

// PUBLIC METHODS =================================================================================

func (self *Checker) And(matcher MatcherFactory, args ...interface{}) *Checker {
	self.negated = false
	if !self.failed {
		return self.run(matcher, args...)
	}
	return self
}

// PRIVATE METHODS ================================================================================

func (self *Checker) run(matcher MatcherFactory, args ...interface{}) *Checker {
	if notMatcher, ok := matcher.(*NotMatcher); ok {
		self.negated = !self.negated
		matcher = notMatcher.inner
	}

	inst := matcher.New(args)
	if self.negated {
		self.failed = !self.Actual.ToNot(inst)
	} else {
		self.failed = !self.Actual.To(inst)
	}
	return self
}
