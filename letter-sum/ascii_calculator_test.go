package letter_sum_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	letterSum "code-katas/letter-sum"
)

func TestLetterSum_CalculateASCIISum(t *testing.T) {
	type casTest struct {
		word string
		sum  int
	}
	casTests := []casTest{
		{"", 0},
		{"a", 1},
		{"z", 26},
		{"cab", 6},
		{"excellent", 100},
		{"microspectrophotometries", 317},
	}

	for _, tt := range casTests {
		t.Run(fmt.Sprintf("Given word \"%s\" we expect %d to be returned as sum", tt.word, tt.sum), func(t *testing.T) {
			actual := letterSum.CalculateASCIISum(tt.word)
			assert.Equal(t, tt.sum, actual)
		})
	}
}
