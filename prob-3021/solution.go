package prob3021

func FlowerGame(n int, m int) int64 {
	isNEven := n%2 == 0
	isMEven := m%2 == 0

	var result int64

	if isNEven {
		if isMEven {
			result = int64((n/2)*(m/2)) * 2
		} else {
			result = int64((n/2)*(m+1)/2 + (n / 2 * (m - 1) / 2))
		}
	} else {
		if isMEven {
			result = int64((n+1)/2*(m/2) + (n-1)/2*(m/2))
		} else {
			result = int64((n+1)/2*(m-1)/2 + (n-1)/2*(m+1)/2)
		}
	}

	return result
}
