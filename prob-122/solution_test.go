package prob122_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-122"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"Test case 1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"Test case 2", []int{1, 2, 3, 4, 5}, 4},
		{"Test case 3", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solution.MaxProfit(tt.prices)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
