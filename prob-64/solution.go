package prob64

import "math"

var cache [][]int

func MinPathSum(grid [][]int) int {
	// base case
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}

	cache = make([][]int, len(grid))
	for i := range cache {
		cache[i] = make([]int, len(grid[0]))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return grid[0][0] + min(calculatePathSum(1, 0, grid), calculatePathSum(0, 1, grid))
}

func calculatePathSum(row, col int, grid [][]int) int {
	if !isValidCell(row, col, grid) {
		return math.MaxInt
	}

	if row == len(grid)-1 && col == len(grid[0])-1 {
		return grid[row][col]
	}

	if cache[row][col] != -1 {
		return cache[row][col]
	}

	cache[row][col] = grid[row][col] + min(calculatePathSum(row+1, col, grid), calculatePathSum(row, col+1, grid))

	return cache[row][col]
}

func isValidCell(row, col int, grid [][]int) bool {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return false
	}
	return true
}
