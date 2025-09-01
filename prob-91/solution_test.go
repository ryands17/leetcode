package prob91_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-91"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
	}

	for _, test := range tests {
		result := solution.NumDecodings(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
