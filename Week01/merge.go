package main

import "sort"
//合并
//1：先合并两个数组，然后在进行排序
//2: 用nums2 的n作为指针，然后从大 到小依次比较

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1 = append(nums1[:m],nums2[:n]...)
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int)  {

	for n>0{

		if m > 0 && nums1[m - 1] > nums2[n-1]{
			nums1[m+n-1] = nums1[m-1]
			m--
		}else{
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}


}