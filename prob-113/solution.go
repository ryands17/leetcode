package prob113

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func PathSum(root *TreeNode, targetSum int) [][]int {
	// bases cases
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return [][]int{{root.Val}}
		}
		return [][]int{}
	}

	res = make([][]int, 0, 10)

	calculate(root, targetSum, nil)

	return res
}

func calculate(node *TreeNode, targetSum int, path []int) {
	if path == nil {
		path = make([]int, 0, 10)
	}

	currentSum := targetSum - node.Val

	path = append(path, node.Val)
	fmt.Println("calculating for", node.Val, path, currentSum)
	if node.Left != nil {
		calculate(node.Left, currentSum, path)
	}
	if node.Right != nil {
		calculate(node.Right, currentSum, path)
	}

	// if it's a leaf node, check if the current sum is 0
	if node.Left == nil && node.Right == nil && currentSum == 0 {
		// make a copy of path and add it to results
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	// pop the last element as it's not needed anymore
	path = path[:len(path)-1]
}
