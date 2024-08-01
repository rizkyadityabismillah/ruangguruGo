package calculator

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"calculator/calculator"
)

var _ = Describe("Calculator", func() {
    Describe("Add", func() {
        It("Should add the base number with adder number", func() {
            calc := calculator.InitCalculator(3)
            calc.Add(2)
            Expect(calc.Result()).To(Equal(5.0))
        })
    })

    Describe("Subtract", func() {
        It("Should subtract the base number with subtractor number", func() {
            calc := calculator.InitCalculator(10)
            calc.Subtract(5)
            Expect(calc.Result()).To(Equal(5.0))
        })
    })

    Describe("Multiply", func() {
        It("Should multiply the base number with multiplier number", func() {
            calc := calculator.InitCalculator(2)
            calc.Multiply(2)
            Expect(calc.Result()).To(Equal(4.0))
        })
    })

    Describe("Divide", func() {
        It("Should return error when the divisor is 0", func() {
            calc := calculator.InitCalculator(2)
            err := calc.Divide(0)
            Expect(err).ToNot(BeNil())
            Expect(err.Error()).To(Equal("cannot divide by zero"))
        })

        It("Should divide the base number with divisor number", func() {
            calc := calculator.InitCalculator(10)
            calc.Divide(2)
            Expect(calc.Result()).To(Equal(5.0))
        })
    })
})
