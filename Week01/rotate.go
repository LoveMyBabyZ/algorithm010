package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7}
	rotate(nums,3)
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums[:len(nums)-k])
	fmt.Println(nums)
	reverse(nums[len(nums)-k:])
	fmt.Println(nums)

	reverse(nums)
	fmt.Println(nums)

}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--

	}
}
