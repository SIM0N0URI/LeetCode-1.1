func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for _, w := range walls {
		grid[w[0]][w[1]] = 1
	}

	for _, g := range guards {
		grid[g[0]][g[1]] = 1
	}

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for _, g := range guards {
		r, c := g[0], g[1]
		for _, dir := range directions {
			dr, dc := dir[0], dir[1]
			nr, nc := r+dr, c+dc

			for nr >= 0 && nr < m && nc >= 0 && nc < n {
				if grid[nr][nc] == 1 {
					break
				}
				grid[nr][nc] = 2
				nr += dr
				nc += dc
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}

	return count
}