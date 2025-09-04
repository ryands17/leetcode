package prob129

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	sum   int
	queue []int
)

func SumNumbers(root *TreeNode) int {
	sum = 0
	queue = []int{}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	sumNumbers(root)
	return sum
}

func sumNumbers(root *TreeNode) {
	queue = append(queue, root.Val)
	if root.Left != nil {
		sumNumbers(root.Left)
	}
	if root.Right != nil {
		sumNumbers(root.Right)
	}
	// this is a leaf node
	if root.Left == nil && root.Right == nil {
		sum += calculateNumber(queue)
	}
	// backtrack
	queue = queue[:len(queue)-1]
}

func calculateNumber(queue []int) int {
	strQueue := make([]string, len(queue))
	for i, num := range queue {
		strQueue[i] = strconv.Itoa(num)
	}
	str := strings.Join(strQueue, "")
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}
