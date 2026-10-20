package main

import "fmt"

// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// or use this method
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}


// STACK(LIFO)
// type stack[T any] struct {
// 	elements []T
// }

func main() {
	// myStack := stack[string]{
	// 	elements: []string{"golang","typescript"},
	// }
	// fmt.Println((myStack))


	// nums := []int{1,2,3}
	names := []string{"golang","typescript"}
	printSlice(names, "john")

}