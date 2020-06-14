package main




func plusOne(digits []int) []int {

	i := len(digits) -1
	for i >=0  {
		//遇到不为9 的值 直接返回
		if digits[i] != 9 {
			digits[i]++
			return digits
		}else{
			digits[i] = 0
		}
		i--
	}

	//全部为9 的情况，在最前端 加1
	return append([]int{1},digits...)
}