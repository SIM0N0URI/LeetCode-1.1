// smallestNumber returns the smallest number of the form 2^k - 1 that is >= n.
func smallestNumber(n int) int {
	var x = 1
	for x < n {
		x = x*2 + 1
	}
	return x
}