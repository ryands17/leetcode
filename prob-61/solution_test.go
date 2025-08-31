package prob61_test

import (
	"testing"

	prob61 "github.com/ryands17/leetcode/prob-61"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		input    *prob61.ListNode
		k        int
		expected *prob61.ListNode
	}{
		{
			input:    createLinkedList([]int{1, 2, 3, 4, 5}),
			k:        2,
			expected: createLinkedList([]int{4, 5, 1, 2, 3}),
		},
		{
			input:    createLinkedList([]int{0, 1, 2}),
			k:        4,
			expected: createLinkedList([]int{2, 0, 1}),
		},
		{
			input:    createLinkedList([]int{}),
			k:        1,
			expected: createLinkedList([]int{}),
		},
		{
			input:    createLinkedList([]int{1, 2, 3}),
			k:        0,
			expected: createLinkedList([]int{1, 2, 3}),
		},
		{
			input:    createLinkedList([]int{1, 2}),
			k:        2,
			expected: createLinkedList([]int{1, 2}),
		},
	}

	for _, test := range tests {
		result := prob61.RotateRight(test.input, test.k)
		if !compareLinkedLists(result, test.expected) {
			t.Errorf("For input %v and k=%d, expected %v but got %v", linkedListToSlice(test.input), test.k, linkedListToSlice(test.expected), linkedListToSlice(result))
		}
	}
}

// Helper function to create a linked list from a slice
func createLinkedList(values []int) *prob61.ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &prob61.ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &prob61.ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to compare two linked lists
func compareLinkedLists(l1, l2 *prob61.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// Helper function to convert a linked list to a slice
func linkedListToSlice(head *prob61.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
