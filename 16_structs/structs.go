package main

import (
	"fmt"
	"time"
)

// order struct

// example of struct embeddings
type customer struct {
	name string
	phone string
}

// composition
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time  // nanosecond precision
	customer
}


// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here....
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
	
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) getAmount() float32 {
// 	return o.amount
// }

func main () {

	// newCustomer := customer{
	// 	name: "jhon",
	// 	phone: "1234567890",
	// }
	newOrder := order{
		id: "1",
		amount: 30,
		status: "received",
		// customer: newCustomer,
		customer: customer{	name: "jhon",
		phone: "1234567890",},
	}
	// fmt.Println(newOrder.customer)
	fmt.Println(newOrder)



	// language := struct {
	// 	name string
	// 	isGood bool
	// }{"golang",true}
	// fmt.Println(language)


	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder)

	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())


	// myOrder.createdAt = time.Now()

	// fmt.Println(myOrder.status)

	// myOrder2 := order {
	// 	id: "2",
	// 	amount: 100,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }


	// myOrder.status = "paid"

	// fmt.Println("order struct", myOrder2)
	// fmt.Println("order struct",myOrder)
}