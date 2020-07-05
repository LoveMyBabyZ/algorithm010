package main

import "fmt"



func numIslands(grid [][]byte) int {
	len1 := len(grid)
	if len1 == 0 {
		return 0
	}
	len2 := len(grid[0])
	res := 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if grid[i][j] == '1' {
				helper(grid,i,j)
				res++
			}
		}
	}
	fmt.Println(grid)
	return res

}

func helper(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		helper( grid,i+1, j,)
	}
	if i-1 >= 0 && grid[i-1][j] == 1 {
		helper( grid,i-1, j,)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
		helper(grid,i, j+1, )
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		helper( grid,i, j-1,)
	}
}
