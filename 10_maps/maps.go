package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map
	// m := make(map[string]string)
	// // setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"
	// get an element
	// fmt.Println(m["name"],m["area"])
	// IMP: if key does not exists in the map then it returns zero value


	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// check map length
	// fmt.Println(len(m))

	// if want to delete from map 
	// delete(m, "price")
	// for remove all map then use clear() func
	// clear(m)
	// fmt.Println(m)


	// for making map i use make and different ways to create map without use of make func
	// m := map[string]int{"price": 40, "phones": 3}
	// // fmt.Println(m)


	//  m := map[string]int{"price": 40, "phones": 3}
	//  v, ok := m["price"]
	//  fmt.Println(v)
	//  if ok {
	// 	fmt.Println("all ok")
	//  } else {
	// 	fmt.Println("not ok")
	//  }


	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}
	// for checking two maps are eqal or not into map using equal not use m1 == m2, just tell true or false
	fmt.Println(maps.Equal(m1,m2))



}