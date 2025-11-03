func maxProduct(nums []int) int64 {
	if len(nums) < 3 {
		return 0
	}

	top1, top2 := 0, 0
	for _, v := range nums {
		n := v
		if n < 0 {
			n = -n
		}
		if n > top1 {
			top2 = top1
			top1 = n
		} else if n > top2 {
			top2 = n
		}
	}

	return int64(top1) * int64(top2) * 100000
}