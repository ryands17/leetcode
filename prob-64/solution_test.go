package prob64_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-64"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 3}, {1, 5}},
			expected: 7,
		},
		{
			input:    [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
	}

	for _, test := range tests {
		result := solution.MinPathSum(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
