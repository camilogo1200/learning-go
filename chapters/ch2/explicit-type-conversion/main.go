package main

import "fmt"

func main() {
	//Go doesn't allow automatic type promotion between variables
	//you must use a type conversion when variable types does not match
	var x = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

	var b byte = 226
	var i int = 30

	var sum3 int = i + int(b)
	var sum4 byte = b + byte(i)
	fmt.Println(sum3, sum4)
}
