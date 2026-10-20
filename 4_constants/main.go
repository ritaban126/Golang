package main

import "fmt"

// const age = 30

func main() {
	// :=
	// const name string = "golang"
	// const age = 30

	// fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
