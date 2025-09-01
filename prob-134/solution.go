package prob134

func CanCompleteCircuit(gas []int, cost []int) int {
	validPosition := 0
	tripCost, totalCost := 0, 0

	for i := range gas {
		tripCost += (gas[i] - cost[i])
		totalCost += (gas[i] - cost[i])

		if tripCost < 0 {
			validPosition = i + 1
			tripCost = 0
		}
	}

	if totalCost < 0 {
		return -1
	}
	return validPosition
}
