package Week06

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	if len(grid[0]) == 1 {
		res := 0
		for i := 0; i < n; i++ {
			res += grid[i][0]
		}
		return res
	}
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		if i == 0 {
			dp[0][0] = grid[0][0]
			continue
		}
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1;i < n;i++{
		for j := 1;j<m;j++{
			dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[n-1][m-1]
}

func min(x,y int)int{
	if x < y {
		return x
	}
	return y
}
