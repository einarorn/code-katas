package game_of_threes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	gameofthrees "code-katas/game-of-threes"
)

func TestGameOfThrees_DivideAndPrintUntilOne(t *testing.T) {
	t.Run("Given we receive the number 100 When we divide by 3 until we got 1 Then we should expect certain output", func(t *testing.T) {
		expected := []int{100, 99, 33, 11, 12, 4, 3, 1}

		actual := gameofthrees.DivideUntilOne(100)

		assert.Equal(t, expected, actual)
	})

	t.Run("Given we receive the number 1 When we divide by 3 until we got 1 Then we should expect only 1", func(t *testing.T) {
		expected := []int{1}

		actual := gameofthrees.DivideUntilOne(1)

		assert.Equal(t, expected, actual)
	})
}
