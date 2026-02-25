package main

import "fmt"

func main() {
	//go naming must be or a letter or an underscore
	var myvar int
	var _myvar string
	var __x string

	fmt.Printf("", myvar, _myvar, __x)

	//the name can contain numbers, underscores, and letters
	//any Unicode character considered a letter or digit is allowed.
	var Ж = "Cyrillic capital letter"
	var א = "alef"
	var ඨ = " Sinhala letter"
	var ꮺ = "Cherokee small letter"
	var 𓌉 = " Egyptian hieroglyph"
	fmt.Printf("", Ж, א, ඨ, ꮺ, 𓌉)

	//Idiomatic Go does not use snake case like use_case or index_counter
	//idiomatic go uses camel case useCase, indexCounter, numberTries
	//the smaller the scope of a variable the smaller the name that's used for it.
	//ex: single letter variable name in for loops
	//ex k,v in a for range loop
}
