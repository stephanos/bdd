package bdd

import "github.com/onsi/gomega/types"

// gomegaMatcher wraps a Result into the gomega.OmegaMatcher interface.
// This way it can easily be passed to Gomega for evaluation.
type gomegaMatcher struct {
	r Result
}

func newGomegaMatcher(result Result) types.GomegaMatcher {
	return &gomegaMatcher{result}
}

func (m *gomegaMatcher) Match(_ interface{}) (success bool, err error) {
	return m.r.Success, m.r.Error
}

func (m *gomegaMatcher) FailureMessage(_ interface{}) string {
	return m.r.FailureMessage
}

func (m *gomegaMatcher) NegatedFailureMessage(_ interface{}) string {
	return m.r.NegatedFailureMessage
}

func resultFromGomega(matcher types.GomegaMatcher, obtained interface{}) Result {
	success, err := matcher.Match(obtained)
	return Result{success, err, matcher.FailureMessage(obtained), matcher.NegatedFailureMessage(obtained)}
}
