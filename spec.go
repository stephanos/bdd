package bdd

import (
	"github.com/onsi/ginkgo"
)

// Describe blocks allow you to organize your specs.
// They can contain any number of BeforeEach, AfterEach, and It blocks.
//
// They are the root of every spec.
func Describe(text string, body func()) bool {
	return ginkgo.Describe(text, body)
}

// With blocks allow you to organize your specs.
// They can contain any number of BeforeEach, AfterEach, and It blocks.
//
// With blocks are usually used inside a Describe block to distinguish
// various scenarios.
func With(text string, body func()) bool {
	return ginkgo.Context(text, body)
}

// When blocks are an alias for With blocks.
func When(text string, body func()) bool {
	return ginkgo.Context(text, body)
}

// It blocks allow you to organize your specs.
// They can not contain other blocks, only assertions.
//
// Normally It blocks are run synchronously. To perform asynchronous tests,
// pass a function that accepts a Done channel. When you do this, you can alsos provide an optional timeout.
func It(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(text, body, timeout...)
}

// BeforeEach blocks are run before It blocks. When multiple BeforeEach blocks
// are defined in nested Describe and Context blocks the outermost
// BeforeEach blocks are run first.
//
// Like It blocks, BeforeEach blocks can be made asynchronous by providing a
// body function that accepts a Done channel
func BeforeEach(body func(), timeout ...float64) bool {
	return ginkgo.BeforeEach(body, timeout...)
}

// AfterEach blocks are run after It blocks. When multiple AfterEach blocks
// are defined in nested Describe and Context blocks the innermost
// AfterEach blocks are run first.
//
// Like It blocks, AfterEach blocks can be made asynchronous by providing a
// body function that accepts a Done channel
func AfterEach(body func(), timeout ...float64) bool {
	return ginkgo.AfterEach(body, timeout...)
}

// BeforeSuite blocks are run just once before any specs are run.
// When running in parallel, each parallel node process will call BeforeSuite.
//
// BeforeSuite blocks can be made asynchronous by providing a body function
// that accepts a Done channel
//
// You may only register *one* BeforeSuite handler per test suite.
func BeforeSuite(body interface{}, timeout ...float64) bool {
	return ginkgo.BeforeSuite(body, timeout...)
}

// AfterSuite blocks are *always* run after all the specs regardless of whether
// specs have passed or failed. Moreover, if an interrupt signal (^C) is received
// it will attempt to run the AfterSuite before exiting.
//
// AfterSuite blocks can be made asynchronous by providing a body function
// that accepts a Done channel
//
// You may only register *one* AfterSuite handler per test suite.
func AfterSuite(body interface{}, timeout ...float64) bool {
	return ginkgo.AfterSuite(body, timeout...)
}
