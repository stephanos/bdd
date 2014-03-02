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

func newMatcherRunner(info *MatcherInfo, expected []interface{}) *MatcherRunner {
	return &MatcherRunner{info, expected}
}

func (mi *MatcherInfo) New(expected []interface{}) Matcher {
	return newMatcherRunner(mi, expected)
}

func (mi *MatcherRunner) Match(actual interface{}) (success bool, message string, err error) {
	if err := mi.validate(mi.expected); err != nil {
		return false, "", err
	}
	return mi.Matcher(actual, mi.expected)
}

func (mi *MatcherInfo) validate(expected []interface{}) error {
	wantParams := mi.Parameters
	if wantParams != -1 {
		haveParams := len(expected)
		if haveParams != wantParams {
			return fmt.Errorf("expected %d parameters, but got %d", wantParams, haveParams)
		}
	}
	return nil
}
