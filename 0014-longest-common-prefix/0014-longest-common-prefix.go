func longestCommonPrefix(strs []string) string {
	pre := strs[0]
	for _, word := range strs[1:] {
		for len(pre) > 0 && (len(word) < len(pre) || word[:len(pre)] != pre) {
			pre = pre[:len(pre)-1]
		}
		if pre == "" {
			return pre
		}
	}
	return pre
}