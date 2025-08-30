package prob45_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-45"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 1, 4, 4}, 2},
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1}, 0},
	}

	for _, test := range tests {
		result := solution.Jump(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
