func isPalindrome(x int) bool {
	RightToLeft := ""
	LeftToRight := strconv.Itoa(x)
	for i := 0; i < len(LeftToRight); i++ {
		RightToLeft = string(LeftToRight[i]) + RightToLeft
	}
	if LeftToRight == RightToLeft {
		return true
	}else{
		return false
	}
}

