package bdd

import (
	"github.com/onsi/ginkgo"
)

func Describe(text string, body func()) bool {
	return ginkgo.Describe(text, body)
}

func With(text string, body func()) bool {
	return ginkgo.Context(text, body)
}

func When(text string, body func()) bool {
	return ginkgo.Context(text, body)
}

func BeforeEach(body func()) bool {
	return ginkgo.BeforeEach(body)
}

func It(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(text, body, timeout...)
}
