package bdd

import (
	"fmt"
)

type MatchFunc func(actual interface{}, expected []interface{}) (success bool, message string, err error)

type Matcher interface {
	Match(actual interface{}) (success bool, message string, err error)
}

type MatcherFactory interface {
	New(expected []interface{}) Matcher
}

type MatcherInfo struct {
	Parameters int
	Matcher    MatchFunc
}

type MatcherRunner struct {
	*MatcherInfo
	expected []interface{}
}

// FACTORY ========================================================================================

func newMatcherRunner(info *MatcherInfo, expected []interface{}) *MatcherRunner {
	return &MatcherRunner{info, expected}
}

// PUBLIC METHODS =================================================================================

func (self *MatcherInfo) New(expected []interface{}) Matcher {
	return newMatcherRunner(self, expected)
}

func (self *MatcherRunner) Match(actual interface{}) (success bool, message string, err error) {
	if err := self.validate(self.expected); err != nil {
		return false, "", err
	}
	return self.Matcher(actual, self.expected)
}

// PRIVATE METHODS ================================================================================

func (self *MatcherInfo) validate(expected []interface{}) error {
	wantParams := self.Parameters
	if wantParams != -1 {
		haveParams := len(expected)
		if haveParams != wantParams {
			return fmt.Errorf("expected %d parameters, but got %d", wantParams, haveParams)
		}
	}
	return nil
}
