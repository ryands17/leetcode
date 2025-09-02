package prob56

import (
	"slices"
)

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	result := make([][]int, 0, len(intervals))
	// sort the interval by the start value
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	result = append(result, intervals[0])
	index := 0
	currentInterval := result[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= currentInterval[1] {
			// merge the interval
			if intervals[i][1] > currentInterval[1] {
				currentInterval[1] = intervals[i][1]
			}
			result[index] = currentInterval
		} else {
			result = append(result, intervals[i])
			index++
			currentInterval = result[index]
		}
	}

	return result
}
