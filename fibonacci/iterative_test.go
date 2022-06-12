package fibonacci_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"code-katas/fibonacci"
)

func TestFibonacci_CalculateByN(t *testing.T) {
	type fibTest struct {
		input    uint
		expected *big.Int
	}
	testArrangement := []fibTest{
		{input: 0, expected: big.NewInt(0)},
		{input: 1, expected: big.NewInt(1)},
		{input: 2, expected: big.NewInt(1)},
		{input: 3, expected: big.NewInt(2)},
		{input: 4, expected: big.NewInt(3)},
		{input: 11, expected: big.NewInt(89)},
		{input: 13, expected: big.NewInt(233)},
		{input: 15, expected: big.NewInt(610)},
		{input: 23, expected: big.NewInt(28657)},
		{input: 28, expected: big.NewInt(317811)},
		{input: 50, expected: big.NewInt(12586269025)},
		{input: 92, expected: big.NewInt(7540113804746346429)},
	}

	for _, ta := range testArrangement {
		t.Run(fmt.Sprintf("Given input %v we expect %v to be returned", ta.input, ta.expected), func(t *testing.T) {
			actual := fibonacci.Fib(ta.input)
			assert.Equal(t, ta.expected, actual)
		})
	}
}
