package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

//send
// func processNum(numChan chan int){
// 	for num := range numChan {
// 		fmt.Println("proccesing number", num)
// 		time.Sleep(time.Second)
// 	}
// 	// fmt.Println("proccesing number", <- numChan)
// }

// received
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// goroutine syschronizer
// func task(done chan bool){
// 	defer func () {done <- true}()
// 	fmt.Println("processing...")
// }

// NOTE -> when you have multiple goroutines then you can use waitgroups but for single use channnels.

// implement queue system for emailchan
// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() {done <- true}()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }


func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func () {
		chan1 <- 10
	} ()
	go func () {
		chan2 <- "pong"
	} ()
	for i := 0; i < 2; i++ {
		select{
		case chan1Val := <- chan1:
		 fmt.Println("received data from chan1",chan1Val)
		case chan2Val := <- chan2:
		 fmt.Println("received data from chan2",chan2Val)
		}
	}



	//  seing he buffer channnel
	// emailChan := make(chan string, 100)  // here 100 is sized
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com",i)  // here %d is placeholder	
	// }
	// fmt.Println("done sending.")
	// // this is important for close
	// close(emailChan)   // for closiong the channel
	// <- done // here ake goroutine ko block karega
	// emailChan <- "1@gmail.com"
	// emailChan <- "2@gmail.com"
	// fmt.Println(<- emailChan)
	// fmt.Println(<- emailChan)


	// goroutine syschronizer func
	// done := make(chan bool)
	// go task(done)
	// <- done // yaha pe ake block ho jayega jabtak data send nhi ho rahe hai tabtak


	// for received func
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <- result // blocking
	// fmt.Println(res)


// for send func
	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// numChan <- 5 // here i sent in channels and that is syntex
	// time.Sleep(time.Second * 2)


	// messageChan := make(chan string)
	// messageChan <- "ping" // blocking
	// // below line means yaha channels ke andar received kar rahe hai or first one send kar rahe hai.
	// msg := <-messageChan
	// fmt.Println(msg)

}