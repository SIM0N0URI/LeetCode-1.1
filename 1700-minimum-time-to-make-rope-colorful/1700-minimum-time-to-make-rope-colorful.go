func minCost(colors string, time []int) int {
	total := 0
	n := len(colors)

	for i := 0; i < n; {
		sum := time[i]
		maxT := time[i]
		j := i + 1

		for j < n && colors[j] == colors[i] {
			sum += time[j]
			if time[j] > maxT {
				maxT = time[j]
			}
			j++
		}

		if j-i > 1 {
			total += sum - maxT
		}
		i = j
	}

	return total
}