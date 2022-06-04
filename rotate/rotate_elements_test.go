package rotate_test

import (
	"code-katas/rotate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameOfThrees_DivideAndPrintUntilOne(t *testing.T) {
	t.Run("Given we receive the slice [1,2,3,4,5,6] When we rotate by 2 Then we should expect [3,4,5,6,1,2] as output", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{3, 4, 5, 6, 1, 2}

		actual := rotate.ShiftLeft(input, 2)

		assert.Equal(t, expected, actual)
	})

	t.Run("Given we receive the slice [1,2] When we rotate by 7 Then we should expect [2,1] as output", func(t *testing.T) {
		input := []int{1, 2}
		expected := []int{2, 1}

		actual := rotate.ShiftLeft(input, 7)

		assert.Equal(t, expected, actual)
	})

	t.Run("Given we receive the slice [1] When we rotate by 3 Then we should expect [1] as output", func(t *testing.T) {
		input := []int{1}
		expected := []int{1}

		actual := rotate.ShiftLeft(input, 3)

		assert.Equal(t, expected, actual)
	})
}
