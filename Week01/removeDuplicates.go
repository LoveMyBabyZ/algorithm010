package main

import "fmt"

func main(){
	nums := []int{1,2,3,3,3,4,5,6}
	removeDuplicates(nums)

}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	left, right := 1, 1
	for right < len(nums) {
		if (nums[right] != nums[right-1]) {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	fmt.Println(nums)
	return left
}
