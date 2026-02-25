package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Zero value for string is the empty string
	//strings are compared for equality using == and for difference != or ordering using <,>,>=,<=
	//strings are concatenated using the + operator
	//Strings in go are immutable, -> you can reassign the value of a string variable,
	//but you cannot change the value of a string that is assigned to it
	var emptyString string
	fmt.Printf("zero value for String : [%s]\n", emptyString)

	var newString string = "aAbBcC"
	var sameString string = "aAbBcC"

	if newString == sameString {
		fmt.Printf("newString [%s] and sameString [%s] are equal \n", newString, sameString)
	} else {
		fmt.Printf("newString [%s] and sameString [%s] are NOT equal \n", newString, sameString)
	}

	var concatenatedString = "Hello "
	concatenatedString = concatenatedString + " - " + newString + " - concatenated String"
	fmt.Println(concatenatedString)

	//Go has a type taht representa single code point. -> The rune type
	//rune type is an alias for the int32 type, just as alias is an alias for uint8
	//if you're referring to a character use -> rune type, not the int32 type

	var myFirstInitial = 'C'
	runeEmoji := '🔥'
	var birthdayMonth = 11
	fmt.Println("My First initial " + string(myFirstInitial))

	var b strings.Builder
	b.WriteString("My First rune emoji: ")
	b.WriteRune(runeEmoji)
	result := b.String()
	fmt.Println(result)
	//rune and an int
	s := strconv.Itoa(birthdayMonth) + string(runeEmoji)
	fmt.Println(s)
	s = fmt.Sprintf("%d%c", birthdayMonth, runeEmoji)
	fmt.Println(s)
}
