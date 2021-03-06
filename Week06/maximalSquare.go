package Week06



func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n+1)
	dp[0] = make([]int, m+1)
	res := 0
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			if matrix[i-1][j-1] == '1'{
				dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
				res = max(res,dp[i][j] * dp[i][j])
			}
		}
	}
	return res
}

func min(a,b,c int)int{
	if a > b{
		if b < c{
			return b
		}
		return c
	}else{
		if a < c{
			return a
		}
		return c
	}
}
func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
