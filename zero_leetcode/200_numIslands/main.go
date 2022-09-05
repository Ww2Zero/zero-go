package _00_numIslands

import (
	"fmt"
	"strconv"
)

func numIslands(grid [][]byte) int {
	res := 0
	fmt.Println(grid)
	fmt.Println("len(grid):", len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				fmt.Println("====" + strconv.Itoa(i) + "====" + strconv.Itoa(j) + "====")
				fmt.Println(grid)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	if !inArea(grid, i, j) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func inArea(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) &&
		j >= 0 && j < len(grid[0])
}
