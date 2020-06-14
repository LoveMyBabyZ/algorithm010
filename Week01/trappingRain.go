package main

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	l := 0
	r := n-1

	maxL := height[l]
	maxR := height[r]

	ans := 0

	for l < r {
		if maxL < maxR {
			ans += maxL - height[l]
			l++
			maxL = max(maxL, height[l])
		} else {
			ans += maxR - height[r]
			r--
			maxR = max(maxR, height[r])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}