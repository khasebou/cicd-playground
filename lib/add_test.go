package lib

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "testing"
)

func TestLibs(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "lib Suite")
}

var _ = Describe("Checking books out of the library", func() {
    When("Adding two numbers", func() {
        It("returns the correct result", func() {
            a := 12
            b := 123
            exp := a + b
            ans := Add(a, b)
            Expect(ans).To(Equal(exp))
        })
    })
})
