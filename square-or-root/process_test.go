package square_or_root_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	square_or_root "code-katas/square-or-root"
)

func TestProcess_ProcessNumber(t *testing.T) {
	type sorTest struct {
		n int // input
		e int // expected
	}
	sorTests := []sorTest{
		{1, 1}, {2, 4}, {3, 9}, {4, 2}, {5, 25}, {6, 36}, {7, 49}, {8, 64}, {9, 3},
	}

	for _, tt := range sorTests {
		t.Run(fmt.Sprintf("Given input %d we expect %d to be returned", tt.n, tt.e), func(t *testing.T) {
			actual := square_or_root.ProcessNumber(tt.n)
			assert.Equal(t, tt.e, actual)
		})
	}
}

func TestProcess_ProcessArray(t *testing.T) {
	input := []int{4, 3, 9, 7, 2, 1}
	expected := []int{2, 9, 3, 49, 4, 1}

	t.Run(fmt.Sprintf("Given input %v we expect %v to be returned", input, expected), func(t *testing.T) {
		actual := square_or_root.ProcessArray(input)

		assert.Equal(t, expected, actual)
	})
}
