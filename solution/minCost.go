package solution

func minCost(costs [][]int) int {
	for i := range costs {
		if i == 0 {
			continue
		}
		costs[i][0] = costs[i][0] + min(costs[i-1][1], costs[i-1][2])
		costs[i][1] = costs[i][1] + min(costs[i-1][0], costs[i-1][2])
		costs[i][2] = costs[i][2] + min(costs[i-1][0], costs[i-1][1])
	}
	return min(min(costs[len(costs)-1][0], costs[len(costs)-1][1]), costs[len(costs)-1][2])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
