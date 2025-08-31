package prob61

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	oldHead := head

	a := head
	b := head
	for range k {
		if b.Next == nil {
			b = head
		} else {
			b = b.Next
		}
	}

	if a == b {
		return head
	}

	for b.Next != nil {
		a = a.Next
		b = b.Next
	}

	newHead := a.Next
	a.Next = nil
	b.Next = oldHead

	return newHead
}
