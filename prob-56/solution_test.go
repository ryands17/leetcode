package prob56_test

import (
	"reflect"
	"testing"

	solution "github.com/ryands17/leetcode/prob-56"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			input:    [][]int{{4, 7}, {1, 4}},
			expected: [][]int{{1, 7}},
		},
	}

	for _, test := range tests {
		result := solution.Merge(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
