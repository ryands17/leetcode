package prob113_test

import (
	"fmt"
	"reflect"
	"testing"

	solution "github.com/ryands17/leetcode/prob-113"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		inputTree []any
		targetSum int
		expected  [][]int
	}{
		{
			inputTree: []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1},
			targetSum: 22,
			expected:  [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			inputTree: []any{1, 2, 3},
			targetSum: 5,
			expected:  [][]int{},
		},
		{
			inputTree: []any{1, 2},
			targetSum: 0,
			expected:  [][]int{},
		},
		{
			inputTree: []any{1, 2},
			targetSum: 1,
			expected:  [][]int{},
		},
	}

	for _, test := range tests {
		root := createTree(test.inputTree)
		result := solution.PathSum(root, test.targetSum)
		fmt.Println("Result:", result)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input tree %v and target sum %d, expected %v but got %v", test.inputTree, test.targetSum, test.expected, result)
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
