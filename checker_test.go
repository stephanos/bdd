package bdd

import "testing"

func TestSuite(t *testing.T) {
	RunSpecs(t, "BDD Suite")
}

var _ = Describe("Suite", func() {

	It("Check", func() {
		Check(1, Equals, 1)
		Check(1, Equals, 1).ButNot(2)
		Check(1, Equals, 1).And(IsLessThan, 2)
		Check("golang", Contains, "go").And("lang").ButNot("gopher")
	})

	It("Assert", func() {
		Assert(1, Equals, 1)
		Assert(1, Equals, 1).ButNot(2)
		Assert(1, Equals, 1).And(IsGreaterThan, 0)
		Assert("golang", Contains, "go").And("lang").ButNot("gopher")
	})

	It("Expect", func() {
		Expect(1, Equals, 1)
		Expect(1, Equals, 1).ButNot(2)
		Expect(1, Equals, 1).And(IsGreaterThanOrEqTo, 1)
		Expect("golang", Contains, "go").And("lang").ButNot("gopher")
	})

	It("Ω", func() {
		Ω(1, Equals, 1)
		Ω(1, Equals, 1).ButNot(2)
		Ω(1, Equals, 1).And(IsLessThanOrEqTo, 1)
		Ω("golang", Contains, "go").And("lang").ButNot("gopher")
	})
})
