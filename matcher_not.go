package bdd

type NotMatcher struct {
	inner MatcherFactory
}

func (nm *NotMatcher) New(expected []interface{}) Matcher {
	return nil // ignore
}

func Not(matcher MatcherFactory) MatcherFactory {
	return &NotMatcher{matcher}
}
