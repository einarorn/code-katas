package cash_machine_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	cashmachine "code-katas/cash-machine"
)

func TestCashMachine_Withdrawal(t *testing.T) {
	type testCase struct {
		input    float64
		expected string
	}

	testCases := []testCase{
		{input: 3.45, expected: "{'£2': 1, '£1': 1, '20p': 2, '5p': 1}"},
		/*{input: 160, expected: "{'£50': 3, '£10': 1}"},
		{input: 0.13, expected: "{'10p': 1, '2p': 1, '1p': 1}"},
		{input: 2149.69, expected: "{'£50': 42, '£20': 2, '£5': 1, '£2': 2, '50p': 1, '10p': 1, '5p': 1, '2p': 2}"},
		{input: 0.99, expected: "{'50p': 1, '20p': 2, '5p': 1, '2p': 2}"},
		{input: 0.01, expected: "{'1p': 1}"},
		{input: 1.16, expected: "{'£1': 1, '10p': 1, '5p': 1, '1p': 1}"},
		{input: 88.88, expected: "{'£50': 1, '£20': 1, '£10': 1, '£5': 1, '£2': 1, '£1': 1, '50p': 1, '20p': 1, '10p': 1, '5p': 1, '2p': 1, '1p': 1}"},
		{input: 0.009, expected: "{}"},
		{input: 0.011, expected: "{'1p': 1}"},*/
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Given input \"%s\" we expect \"%v\" to be returned", strconv.FormatFloat(tt.input, 'f', -1, 64), tt.expected), func(t *testing.T) {
			actual, err := cashmachine.BreakIntoChange(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	t.Run("Given input is less than zero we return an error", func(t *testing.T) {
		actual, err := cashmachine.BreakIntoChange(-0.001)
		assert.Error(t, err)
		assert.Equal(t, "", actual)
	})
}
