package main

import "fmt"

func main() {
	var x int32 = 30
	pointer := &x

	//&address operator
	fmt.Printf(" Pointer = {%v} - value ={%v}\n", pointer, *pointer)
}
