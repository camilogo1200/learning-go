package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//if code block is very similar to other languages if control flow.
	//however it allows you or is commonly used to declare if statement variables

	//if and else
	if n := rand.N(100); n == 0 {
		fmt.Println("n == 0")
	} else if n > 5 {

	}

	//Scoping a variable to an if statement
	if n := rand.N(100); n == 0 {
		fmt.Printf("n == 0. n = {%d}", n)
	} else {
		fmt.Printf("n ! = 0. n = {%d}", n)
	}

	//fmt.Println(n) //n only scoped to the if statement block

}
