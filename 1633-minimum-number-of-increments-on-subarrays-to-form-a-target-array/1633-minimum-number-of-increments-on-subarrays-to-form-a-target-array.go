// minNumberOperations calculates the minimum number of increment operations needed
func minNumberOperations(target []int) int {
	if len(target) == 0 {
		return 0
	}
	opera := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			opera += target[i] - target[i-1]
		}
	}
	return opera
}