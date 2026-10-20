package main

import "fmt"

// enumerated types

type OrderStatus int

// iota just for integer you can write as a string in type then in iota write "received" then others write same 
const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to",status)
}

func main() {
	changeOrderStatus(Received)

}