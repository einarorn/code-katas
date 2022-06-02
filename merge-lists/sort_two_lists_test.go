package merge_lists_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	mergelists "code-katas/merge-lists"
)

func TestMergeLists_SortTwoLists(t *testing.T) {
	type stlTest struct {
		input1   []int
		input2   []int
		expected []int
	}
	testArrangement := []stlTest{
		{input1: []int{1, 4, 6}, input2: []int{2, 3, 5}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input1: []int{2, 3, 5}, input2: []int{1, 4, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input1: []int{1, 2, 3}, input2: []int{4, 5, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input1: []int{4, 5, 6}, input2: []int{1, 2, 3}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input1: []int{1, 1}, input2: []int{1, 2}, expected: []int{1, 1, 1, 2}},
		{input1: []int{1, 2}, input2: []int{1, 1}, expected: []int{1, 1, 1, 2}},
		{input1: []int{2}, input2: []int{1}, expected: []int{1, 2}},
		{input1: []int{1}, input2: []int{2}, expected: []int{1, 2}},
	}

	for _, ta := range testArrangement {
		t.Run(fmt.Sprintf("Given input1 %v and input2 %v we expect %v to be returned", ta.input1, ta.input2, ta.expected), func(t *testing.T) {
			actual := mergelists.SortTwoLists(ta.input1, ta.input2)
			assert.Equal(t, ta.expected, actual)
		})
	}
}
