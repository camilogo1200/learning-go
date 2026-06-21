package main

import "fmt"

//build  function called `prefixer` that has an input parameter as string, and returns a function that receives a string,
//and uses the string paramater as a prefix whe you call the function

func main() {

	helloPrefix := prefixer("Hello ")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Caroline"))

}

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		str := prefix + s
		return str
	}

}
