package main


import (
	"fmt"
	// "time"
)

func main() {
	// simple switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
    // case 3:
	// 	fmt.Println("three")
	// // default:
	// // 	fmt.Println("others") default is optional
	// }

     
	// multiple contidion switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }


	// type switch. here whoAmI is a function and interface{} ke baad jo bracket hai wo batate hai ki i ki value kuch hi ho sakte hai.
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int: 
		  fmt.Println("It's an integer")
		case string: 
		  fmt.Println("It's a string")
		case bool: 
		  fmt.Println("It's a boolean")
		default: 
		  fmt.Println("others",t)
		}
	}

	whoAmI(50.3) // function call and above you can remove t and from default also. 
}
