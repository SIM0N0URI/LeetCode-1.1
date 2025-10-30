// minimumTotal computes the minimum path sum from top to bottom in a triangle.
// It uses bottom-up dynamic programming, updating a single array in place.
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	acc := append([]int(nil), triangle[len(triangle)-1]...)
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			acc[j] = triangle[i][j] + min(acc[j], acc[j+1])
		}
	}
	return acc[0]
}