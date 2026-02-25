package main

import "fmt"

func main() {
	//var versus :=
	//verbose way
	var x int = 10
	//if the right side of the = is the expected type of your variable you can remove the type from the left side
	//since the default type for a number is an int
	var myInt = 20
	//if you want to declare a variable and assign it the zero value you remove the right side and the =
	var defaultInt int
	//you can declare multiple variable at once
	var w, y, z int = 10, 20, 30
	//you can declare all zero values of the same type
	var a, b, c int
	//or the different types
	var i, j = 10, "helloWorld"
	// If you're declaring multiple variables at once. you can wrap them in a declaration list
	var (
		t       int
		r             = 20
		e       int64 = 234567890123
		myHello       = " Hello World"
		f, g    string
		ii, jj  float32
	)
	fmt.Printf("\n", x, myInt, defaultInt, w, y, z, a, b, c, i, j, t, r, e, myHello, f, g, ii, jj)

	//Go supports also short declaration
	//use the := operator to replace the var declaration that uses type inference

	var gg = 10
	kk := 10
	//declare multiple variables at once using :=
	l, m := 15, "myString"
	//short declaration has one limitation.
	//if you are declaring a variable at the package level, you must use var declaration, because := is illegal outside of functions
	fmt.Printf("\n", gg, kk, l, m)
}
