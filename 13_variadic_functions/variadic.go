package main

import "fmt"

func sum(nums ...int)int{
	// for any value add you can use ...interface{}
	total := 0
	for _, num := range nums{
		total = total + num
	}
	return total
}

func main() {
	nums := []int{3,4,5,5}
	result := sum(nums...)
	fmt.Println(result)
}