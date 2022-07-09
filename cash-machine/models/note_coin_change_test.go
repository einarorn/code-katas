package models_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"code-katas/cash-machine/models"
)

func TestModels_NotesAndCoinsToString(t *testing.T) {
	type testCase struct {
		input    models.ATMChange
		expected string
	}

	testCases := []testCase{
		{
			input:    []models.NoteCoinChange{{Amount: "£2", Count: 2}, {Amount: "£1", Count: 1}, {Amount: "20p", Count: 3}, {Amount: "5p", Count: 1}},
			expected: "{'£2': 2, '£1': 1, '20p': 3, '5p': 1}",
		},
		{
			input:    []models.NoteCoinChange{},
			expected: "{}",
		},
		{
			input: []models.NoteCoinChange{
				{Amount: "£50", Count: 1},
				{Amount: "£20", Count: 1},
				{Amount: "£10", Count: 1},
				{Amount: "£5", Count: 1},
				{Amount: "£2", Count: 1},
				{Amount: "£1", Count: 1},
				{Amount: "50p", Count: 1},
				{Amount: "20p", Count: 1},
				{Amount: "10p", Count: 1},
				{Amount: "5p", Count: 1},
				{Amount: "2p", Count: 1},
				{Amount: "1p", Count: 1},
			},
			expected: "{'£50': 1, '£20': 1, '£10': 1, '£5': 1, '£2': 1, '£1': 1, '50p': 1, '20p': 1, '10p': 1, '5p': 1, '2p': 1, '1p': 1}",
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Given notes and coins change \"%v\" we expect convertion to string to be \"%s\"", tt.input, tt.expected), func(t *testing.T) {
			fmt.Println(tt.expected)
			actual := tt.input.NotesAndCoinsToString()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
