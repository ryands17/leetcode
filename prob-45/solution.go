package prob45

import (
	"math"
	"slices"
)

func Jump(nums []int) int {
	l := len(nums)

	if l <= 1 {
		return 0
	}

	table := make([]int, l-1)
	lastIndex := l - 1
	for i := l - 2; i >= 0; i-- {
		if nums[i] == 0 {
			table[i] = math.MaxInt32
		} else if nums[i]+i >= lastIndex {
			table[i] = 1
		} else {
			table[i] = slices.Min(table[i+1:min(len(table), i+nums[i]+1)]) + 1
		}
	}

	return table[0]
}
