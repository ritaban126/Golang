package main

import "fmt"


type paymenter interface{
	pay(amount float32) 
	// refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}

// Open close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)

}

// this two is concret implementation hai razorpay ka 
type razorpay struct{}
func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay",amount)
}

// this two is concret implementation hai stripe ka 
// type stripe struct{}
// func (s stripe) pay(amout float32) {
// 	fmt.Println("making payment using stripe",amout)
// }

// type fakepayment struct{}
// func (f fakepayment) pay(amount float32){
// 	fmt.Println("making payment using fake gateway for testing purpose")
// }

func main() { 
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
	// newPayment := payment{}
	// newPayment.makePayment(100)

}


// book name "clean architecture"