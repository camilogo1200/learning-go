package main

import "fmt"

/**
1. Declaring and calling functions
*/

func main() {
	//Simulating Named and Optional parameters
	MyFunc(MyFuncOpts{
		LastName: "Patek",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Cristhian",
		LastName:  "Smith",
	})

	/*
		Go supports variadic parameters
		Variadic Input Parameters and Slices
		Variadic parameter must be the last or the only one in the input parameter list
		You indicate it with three dots (...) before the type
	*/

	sum, _ := addition(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("sum = {%d}\n", sum)
	sum, _ = addition(1)
	fmt.Printf("sum = {%d}\n", sum)
	sum, _ = addition()
	fmt.Printf("sum = {%d}\n", sum)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum, _ = addition(slice...)
	fmt.Printf("sum = {%d}\n", sum)

	//you can declare a function variable -> functions are values
	var myfuncVariable func(...int) (int, error)
	myfuncVariable = addition
	myfuncVariable()
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	fmt.Printf("opts = {%v}\n", opts)
}

func addition(vals ...int) (int, error) {
	sum := 0
	for _, value := range vals {
		sum += value
	}
	return sum, nil
}
