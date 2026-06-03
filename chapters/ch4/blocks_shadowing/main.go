package main

import "fmt"

func main() {
	//block scope is enforced in golang
	x := 33
	fmt.Printf("x = {%d}\n", x)
	if x > 5 {
		x, y := 50, 70
		fmt.Printf("x = {%d}, y = {%d}\n", x, y)
	}
	fmt.Printf("x = {%d}\n", x)
}
