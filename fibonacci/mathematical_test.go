package fibonacci_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"code-katas/fibonacci"
)

func TestFibonacci_CalculateByN(t *testing.T) {
	type fibTest struct {
		input    int
		expected uint64
	}
	testArrangement := []fibTest{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 2, expected: 1},
		{input: 3, expected: 2},
		{input: 4, expected: 3},
		{input: 11, expected: 89},
		{input: 13, expected: 233},
		{input: 15, expected: 610},
		{input: 23, expected: 28657},
		{input: 28, expected: 317811},
	}

	for _, ta := range testArrangement {
		t.Run(fmt.Sprintf("Given input %v we expect %v to be returned", ta.input, ta.expected), func(t *testing.T) {
			actual, err := fibonacci.CalculateByN(ta.input)
			assert.NoError(t, err)
			assert.Equal(t, ta.expected, actual)
		})
	}

	t.Run("Given input calculates a number bigger than float64 we expect an error to be returned", func(t *testing.T) {
		_, err := fibonacci.CalculateByN(1050)
		assert.Error(t, err)
	})
}
