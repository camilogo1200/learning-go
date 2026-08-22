package main

import "fmt"

func main() {
	//sometimes you need to  store a value of any type, Go uses empty interfaces
	var i interface{}
	// any is an alias for interface{}
	//var i any // new code should be on any Go > 1.18

	i = 20
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	i = struct {
		FirstName string
		LastName  string
	}{"bob", "sponge"}
	fmt.Println(i)
}
