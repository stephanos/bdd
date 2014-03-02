package bdd

import (
	"github.com/onsi/gomega"
)

// IsAssignableTo succeeds if actual is assignable to the type of expected.
// It will return an error when one of the values is nil.
//
//        Check(0, IsAssignableTo, 0)         // Same values
//        Check(5, IsAssignableTo, -1)        // different values same type
//        Check("foo", IsAssignableTo, "bar") // different values same type
func IsAssignableTo(expected interface{}) gomega.OmegaMatcher {
	return gomega.BeAssignableToTypeOf(expected)
}
