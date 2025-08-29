package prob3021_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-3021"
)

func TestFlowerGame(t *testing.T) {
	tests := []struct {
		n, m   int
		expect int64
	}{
		{3, 2, 3},
		{3, 3, 4},
		{4, 4, 8},
		{1, 1, 0},
		{1, 5, 2},
	}

	for _, test := range tests {
		result := solution.FlowerGame(test.n, test.m)
		if result != test.expect {
			t.Errorf("FlowerGame(%d, %d) = %d; want %d", test.n, test.m, result, test.expect)
		}
	}
}
