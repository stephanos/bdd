package bdd

type notMatcher struct {
	inner Matcher
}

func (m *notMatcher) Apply(actual interface{}, args []interface{}) (r Result) {
	r = m.inner.Apply(actual, args)
	r.Success = !r.Success
	r.FailureMessage, r.NegatedFailureMessage = r.NegatedFailureMessage, r.FailureMessage
	return
}

func (m *notMatcher) Name() (n string) {
	if nm, ok := m.inner.(namedMatcher); ok {
		n = "!" + nm.Name()
	}
	return
}

// Not negates a matcher.
func Not(m Matcher) Matcher {
	return &notMatcher{m}
}
