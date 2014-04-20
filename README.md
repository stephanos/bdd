bdd [![Build Status](https://secure.travis-ci.org/101loops/bdd.png)](https://travis-ci.org/101loops/bdd) [![Coverage Status](https://coveralls.io/repos/101loops/bdd/badge.png)](https://coveralls.io/r/101loops/bdd) [![GoDoc](https://camo.githubusercontent.com/6bae67c5189d085c05271a127da5a4bbb1e8eb2c/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f736d61727479737472656574732f676f636f6e7665793f7374617475732e706e67)](http://godoc.org/github.com/101loops/bdd)
======

Go package for writing BDD-style tests.

### Installation
`go get github.com/101loops/bdd`

### Example
```go
package bdd

import . "github.com/101loops/bdd"

var _ = Describe("User Service", func() {
	It("loads users by domain name", func() {
		users, err := service.loadUsersByDomain("acme.com")
		Check(err, IsNil)
		Check(users, HasLen, 2)
		Check(users[0].IsActive, IsTrue)
		Check(users[0].FirstName, Equals, "Roger")
		Check(users[0].Bio, Contains, "ACME").And("Roger").ButNot("fired")
	})
})
```

### Documentation
http://godoc.org/github.com/101loops/bdd

### Credit
Uses [Ginkgo](http://onsi.github.io/ginkgo/) and [Gomga](http://onsi.github.io/gomega/) internally.

### License
MIT (see LICENSE).
