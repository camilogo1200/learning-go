package main

import (
	"fmt"
	"io"
)

type MyInt int

func main() {
	//Type Assertions names the concrete type that implemented the interface, or names another interface that is also implemented by the concrete type whose value is stored in the interface
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) //if this assertion is wrong this code produces a panic
	fmt.Println(i2 + 1)
	//producing a panic
	//i3 := i.(string)
	//fmt.Println(i3)

	// you can avoid this situation using the comma syntax idiom
	i4, ok := i.(string)
	if !ok {
		fmt.Println("wrong type i4 not string")
	} else {
		fmt.Println("i4 is string ", i4)
	}

	// when an interface could be one of multiple possible types, use a type switch

	switch j := i.(type) {
	case nil:
		//i is nil, type of j is any
		fmt.Println("i is nil")
	case int:
		fmt.Println("j is of type int", j)
	case MyInt:
		fmt.Println("j is of type MyInt")
	case io.Reader:
		fmt.Println("j is of type io.Reader")
	case string:
		fmt.Println("j is of type string")
	case bool, rune:
		fmt.Println("j is of type bool or rune, so j is of type any")
	default:
		fmt.Println("No idea what is i. So, j is of type any")
	}

}
