package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{6,7,8}

	// for i :=0; i<len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for i, nums := range nums {
	// 	// sum = sum + nums
	// 	fmt.Println(nums,i)
	// }
	// // fmt.Println(sum)


	// m := map[string]string{"fname":"jhon","lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k,v)
	// }


	// u also can use range for string
	// for this programme i got unicode code point rune for 0,1,2-5.
	// and here i is not index its a starting byte of rune
	// 255 -> 1 byte
	for i, c := range "golang" {
		// fmt.Println(i,c)
		fmt.Println(i,string(c))
	}
}