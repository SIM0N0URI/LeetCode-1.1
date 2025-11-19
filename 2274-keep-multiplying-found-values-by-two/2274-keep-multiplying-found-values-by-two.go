func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	res := 0
	for _, num := range nums {
		if num == original || res == num {
			res = num * 2
		}
	}
	if res == 0 {
		return original
	}
	return res
}