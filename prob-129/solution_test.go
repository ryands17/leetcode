package prob129_test

import (
	"testing"

	solution "github.com/ryands17/leetcode/prob-129"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		inputTree []any
		expected  int
	}{
		{
			inputTree: []any{1, 2, 3},
			expected:  25,
		},
		{
			inputTree: []any{4, 9, 0, 5, 1},
			expected:  1026,
		},
	}

	for _, test := range tests {
		root := createTree(test.inputTree)
		result := solution.SumNumbers(root)
		if result != test.expected {
			t.Errorf("For input tree %v, expected %d but got %d", test.inputTree, test.expected, result)
		}
	}
}

// Helper function to create a binary tree from a slice
func createTree(arr []any) *solution.TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	// Create root node
	root := &solution.TreeNode{Val: arr[0].(int)}

	// Use a queue to track nodes that need children assigned
	queue := []*solution.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		current := queue[0]
		queue = queue[1:] // dequeue

		// Add left child
		if i < len(arr) {
			if arr[i] != nil {
				current.Left = &solution.TreeNode{Val: arr[i].(int)}
				queue = append(queue, current.Left)
			}
			i++
		}

		// Add right child
		if i < len(arr) {
			if arr[i] != nil {
				current.Right = &solution.TreeNode{Val: arr[i].(int)}
				queue = append(queue, current.Right)
			}
			i++
		}
	}

	return root
}
