package bdd

import "github.com/onsi/gomega"

// IsClosed succeeds if actual is a closed channel.
//
// In order to check whether or not the channel is closed, it is always necessary
// to read from the channel. You should keep this in mind if you wish to make
// subsequent assertions about values coming down the channel.
//
// Also, if you are testing that a *buffered* channel is closed you must first
// read all values out of the channel before asserting that it is closed.
var IsClosed = &matcher{
	name: "IsClosed",
	apply: func(actual interface{}, _ []interface{}) Result {
		return resultFromGomega(gomega.BeClosed(), actual)
	},
}
