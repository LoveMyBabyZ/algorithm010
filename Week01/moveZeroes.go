package main

import "fmt"

func main(){
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}



//通过两个指针进行遍历 ，
func moveZeroes(nums []int)  {
	i,j := 0,1
	for j = 1; j < len(nums); j++ {
		fmt.Println(nums[j])
		if nums[i] == 0 && nums[j] != 0{
			nums[i],nums[j] = nums[j],nums[i]
			i++
		}else if nums[i] != 0 {
			i++
		}
	}
}