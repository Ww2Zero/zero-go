package _020_numEnclaves

import (
	"fmt"
	"strconv"
)

func numEnclaves(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}
	for i := 1; i < n-1; i++ {
		dfs(grid, 0, i)
		dfs(grid, m-1, i)
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	fmt.Println("====" + strconv.Itoa(i) + "====" + strconv.Itoa(j) + "====")
	if !inArea(grid, i, j) {
		return
	}
	// flood fill 判断是否已经淹没岛屿
	if grid[i][j] == 0 {
		return
	}
	// flood fill 淹没岛屿
	grid[i][j] = 0
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func inArea(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
