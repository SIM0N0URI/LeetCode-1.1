func getSneakyNumbers(nums []int) []int {
	slc1 := []int{}
	mappos := make(map[int]bool)
	for _, n := range nums {
		if mappos[n] == true {
			slc1 = append(slc1, n)
			continue
		}
		mappos[n] = true
	}
	return slc1
}
