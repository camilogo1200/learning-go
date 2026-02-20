package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
	//Literals -> explicitly specified number, character, or string

	//integer literals -> base 10, but you can specify other formats

	var integer int32 = 1_234_567
	var octal int32 = 0o_1234
	var binary int32 = 0b_101
	var hexadecimal int32 = 0x_AA

	fmt.Printf("Integer Literals : \n integer: %d, octal: %d, binary:%d, hex:%d \n\n", integer, octal, binary, hexadecimal)

	//floating-point literals
	var avogadro float64 = 6.03e23
	var coulomb float64 = 6.24e18

	fmt.Printf("floating-point Literals : \n Avogadro: %f, Coulomb: %f\n\n ", avogadro, coulomb)

	//rune literals -> represents a character and is surrounded by single quotes
	//Go double quotes and single quotes are not interchangeable
	//rune literals can be written as single Unicode characters -> 'a'
	//8-bit octal numbers -> ('\141')
	//8-bit hexadecimal numbers -> ('\x61')
	//16-bit hexadecimal numbers -> ('\u0061')
	//32-bit Unicode numbers -> ('\U00000061')
	//there are some rune literals -> newline ('\n'), tab ('\t'), single quote ('\''), and backslash('\\')
	var character rune = 'c'
	var dollar rune = '$'
	var unicode rune = '\U00000061'
	fmt.Printf("rune Literals (Println): character: %v, dollar: %v, unicode : %v\n\n", character, dollar, unicode)

	//If you iterate over a string using a standard
	//for i := 0 byte-loop and print using %c, you will shatter multi-byte characters and print garbage
	// A string containing multi-byte Unicode characters
	str := "日本\x80語"

	// The range keyword automatically parses UTF-8 into runes
	for pos, char := range str {
		// Using %#U to debug the exact memory decoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)[9]
	}

	//format verb:
	// %c  -> Prints the raw Unicode character ( most common format verb for standard text output)
	// %q  -> prints the  character safely escaped as a single quote rune constant ( logging or debugging)
	// %U  -> Prints the standard Unicode code point identifier
	// %#U -> Prints both the Unicode and the quoted character if it is printable
	fmt.Printf("rune Literals (Printf): character:%c, dollar: %c, unicode : %c\n\n", character, dollar, unicode)
	fmt.Printf("rune Literals (Printf): character:%q, dollar: %q, unicode : %q\n\n", character, dollar, unicode)
	fmt.Printf("rune Literals (Printf): character:%U, dollar: %U, unicode : %U\n\n", character, dollar, unicode)
	fmt.Printf("rune Literals (Printf): character:%#U, dollar: %#U, unicode : %#U\n\n", character, dollar, unicode)

	//string literals -> use double quotes to create  interpreted string literals
	var greeting string = "Greetings and Salutations"
	fmt.Printf("string literals:\nGreeting: %v\n", greeting)
	//raw string literal -> `
	var rawString string = `
	Greetings "raw" 
	string "literal"
	`
	fmt.Printf("string literals:\nrawString: %v\n", rawString)
}
