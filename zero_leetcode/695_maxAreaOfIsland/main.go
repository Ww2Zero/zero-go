package _95_maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				tmpArea := areaDfs(grid, i, j)
				res = max(res, tmpArea)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func areaDfs(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2

	return 1 +
		areaDfs(grid, r-1, c) +
		areaDfs(grid, r+1, c) +
		areaDfs(grid, r, c-1) +
		areaDfs(grid, r, c+1)

}
func inArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
