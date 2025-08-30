package prob55

func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var maxPosition int
	for i, num := range nums {
		if i > maxPosition {
			return false
		}
		maxPosition = max(i+num, maxPosition)
	}

	return true
}
