package main

import (
	"fmt"
)

func main() {
	// for is the only word for loops in Golang the only keyword
	//there are four ways to implement for :
	// 1. c-Style for: for i:=1; ...
	// 2. A condition only for
	// 3. And infinite Loop
	// 4. For range

	//1. traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()

	// you can do the initialization of the loop variable outside, and the increment inside the loop, depending on your conditions
	//2. Condition only for
	i := 0
	for i < 10 {
		fmt.Printf(" %d", i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
	fmt.Println()

	// 3. Infinite for Loop
	//for {
	//	fmt.Println("Hi")
	//}

	//`break` and ` continue`
	//the ` break` statement exits the loop immediately
	//you can use break when you need it not only on the infinite loop

	//continue keyboard skips the rest of the loop for the iteration and proceeds to the next iteration
	//technically you don't need the continue keyword, writing the code in another way, but this keyword helps in different scenarios

	//4. FOR - RANGE statement
	//you can use range for string, arrays, slices, maps and channels
	//slices
	evenVals := []int{2, 4, 6, 8, 10}
	for index, value := range evenVals {
		fmt.Printf("index = {%d} - value = {%d}\n", index, value)
	}

	//maps
	m := map[string]int{
		"a": 1,
		"b": 3,
		"c": 6,
	}

	for index, value := range m {
		fmt.Printf(" [ index = {%s}, value = {%d} ]\n", index, value)
	}

	//Strings

	samples := []string{"Hello", "World", "!"}

	for index, sample := range samples {
		fmt.Printf(" [ index = {%d}, sample = {%s} ]\n", index, sample)

		for i := 0; i < len(sample); i++ {
			fmt.Printf(" [ index = {%d}, sample = {%c]\n", i, sample[i])
		}

		for internalIndex, character := range sample {
			fmt.Printf(" [ index = {%d}, character = {%d}, string(character) = {%s}]\n", internalIndex, character, string(character))
		}
	}

}
