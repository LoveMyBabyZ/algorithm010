package Week06

func numDecodings(s string)int{
	if len(s) == 0 || s == "0" || s[0] == '0'{
		return 0
	}
	n := len(s)
	dp := make([]int,n)
	dp[0] = 1
	for i := 1;i< n;i++{
		if s[i] != '0'{
			dp[i] = dp[i-1]
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6'){
			if i >= 2 {
				dp[i] += dp[i-2]
			}else{
				dp[i]++
			}
		}
	}
	return dp[n-1]
}
