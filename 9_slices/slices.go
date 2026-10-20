package main

import (
	"fmt"
	// "slices"
)

//slice -> dynamic
// most used construct in go
// + usefull method
func main() {
	// uninitialized slice is nill(means null for go use nill)
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))


	// var nums = make([]int,0,5)
	// capacity -> maximum numbers of elements can fit.
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)
	// for adding a number
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// var nums = make([]int,0,5) // first append line in next line then check length and then copy
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy function in slice
	// copy(nums2,nums)
	// fmt.Println(nums,nums2)


	// slice operator
	// var nums = []int{1,2,3}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:2])


	// slices
	// var nums1 = []int{1,2,3}
	// var nums2 = []int{1,2,4}
	// fmt.Println(slices.Equal(nums1,nums2))


	// 2D slices
	var nums = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)


}