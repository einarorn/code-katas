package cash_machine_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	cashmachine "code-katas/cash-machine"
	"code-katas/cash-machine/models"
)

func TestCashMachine_WithdrawalGBP(t *testing.T) {
	type testCase struct {
		input    float64
		expected models.ATMChange
	}

	testCases := []testCase{
		{input: 3.45, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "£2", Count: 1},
			models.NoteCoinChange{Amount: "£1", Count: 1},
			models.NoteCoinChange{Amount: "20p", Count: 2},
			models.NoteCoinChange{Amount: "5p", Count: 1},
		}},
		{input: 160, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "£50", Count: 3},
			models.NoteCoinChange{Amount: "£10", Count: 1},
		}},
		{input: 0.13, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "10p", Count: 1},
			models.NoteCoinChange{Amount: "2p", Count: 1},
			models.NoteCoinChange{Amount: "1p", Count: 1},
		}},
		{input: 2149.69, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "£50", Count: 42},
			models.NoteCoinChange{Amount: "£20", Count: 2},
			models.NoteCoinChange{Amount: "£5", Count: 1},
			models.NoteCoinChange{Amount: "£2", Count: 2},
			models.NoteCoinChange{Amount: "50p", Count: 1},
			models.NoteCoinChange{Amount: "10p", Count: 1},
			models.NoteCoinChange{Amount: "5p", Count: 1},
			models.NoteCoinChange{Amount: "2p", Count: 2},
		}},
		{input: 0.99, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "50p", Count: 1},
			models.NoteCoinChange{Amount: "20p", Count: 2},
			models.NoteCoinChange{Amount: "5p", Count: 1},
			models.NoteCoinChange{Amount: "2p", Count: 2},
		}},
		{input: 0.01, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "1p", Count: 1},
		}},
		{input: 1.16, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "£1", Count: 1},
			models.NoteCoinChange{Amount: "10p", Count: 1},
			models.NoteCoinChange{Amount: "5p", Count: 1},
			models.NoteCoinChange{Amount: "1p", Count: 1},
		}},
		{input: 88.88, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "£50", Count: 1},
			models.NoteCoinChange{Amount: "£20", Count: 1},
			models.NoteCoinChange{Amount: "£10", Count: 1},
			models.NoteCoinChange{Amount: "£5", Count: 1},
			models.NoteCoinChange{Amount: "£2", Count: 1},
			models.NoteCoinChange{Amount: "£1", Count: 1},
			models.NoteCoinChange{Amount: "50p", Count: 1},
			models.NoteCoinChange{Amount: "20p", Count: 1},
			models.NoteCoinChange{Amount: "10p", Count: 1},
			models.NoteCoinChange{Amount: "5p", Count: 1},
			models.NoteCoinChange{Amount: "2p", Count: 1},
			models.NoteCoinChange{Amount: "1p", Count: 1},
		}},
		{input: 0.009, expected: models.ATMChange{}},
		{input: 0.011, expected: models.ATMChange{
			models.NoteCoinChange{Amount: "1p", Count: 1},
		}},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Given we recieve order to withdraw amount of \"%s\" we expect to receive \"%v\" as change", strconv.FormatFloat(tt.input, 'f', -1, 64), tt.expected), func(t *testing.T) {
			cm := cashmachine.New(cashmachine.InitializeGPB())
			actual, err := cm.BreakIntoChange(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
