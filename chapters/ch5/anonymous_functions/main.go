package main

import (
	"fmt"
	"math"
)

// You can declare var with anonymous functions in package scope
var (
	add = func(a float32, b float32) (float32, error) { return a + b, nil }
	sub = func(a float32, b float32) (float32, error) { return a - b, nil }
	mul = func(a float32, b float32) (float32, error) { return a * b, nil }
	mod = func(a float32, b float32) (float32, error) {
		return float32(math.Mod(float64(a), float64(b))), nil
	}
)

func main() {
	//you not only can assign functions top variables but also
	// define new functions within a function and assign them to a variables
	// this approach is useful in `defer` statements and launching goroutines
	f := func(j int) {
		fmt.Printf("AF -> number : [%d]\n", j)
	}

	for i := 0; i < 9; i++ {
		f(i)
	}

	//same code or equivalent
	for i := 0; i < 9; i++ {
		func(j int) {
			fmt.Printf("inner call func -> number : [%d]\n", j)
		}(i)
	}

	//package var definition
	result, ok := add(3.20, 4.00)
	if ok != nil {
		fmt.Println("error")
	}

	fmt.Printf("package function call func -> add : [%.2f]\n", result)

	//NOTE: anonymous functions should be avoided on hot path, the right way to implement is as struct type function to avoid extra allocations
	//Closure are functions declared inside functions
	//functions declared inside functions are able to access and modify variables declared in the outer function.
}
