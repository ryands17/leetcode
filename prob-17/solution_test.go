package prob17_test

import (
	"reflect"
	"testing"

	solution "github.com/ryands17/leetcode/prob-17"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := solution.LetterCombinations(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}
