package prob134_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-134"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
	}

	for _, test := range tests {
		result := solution.CanCompleteCircuit(test.gas, test.cost)
		if result != test.expected {
			t.Errorf("For gas %v and cost %v, expected %d but got %d", test.gas, test.cost, test.expected, result)
		}
	}
}
