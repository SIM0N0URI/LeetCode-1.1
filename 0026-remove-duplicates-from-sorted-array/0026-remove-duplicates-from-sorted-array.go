func removeDuplicates(nums []int) int {
	var unique int
	for i := 0; i < len(nums); i++ {
		duplicate := false
		for j := 0; j < unique; j++ {
			if nums[i] == nums[j] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			nums[unique] = nums[i]
			unique++
		}
	}
	return unique
}