package prob55_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-55"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{2, 3, 1, 4, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
	}

	for _, test := range tests {
		result := solution.CanJump(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
