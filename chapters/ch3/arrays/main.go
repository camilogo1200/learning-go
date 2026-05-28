package main

import "fmt"

func main() {
	// Arrays are rarely used in Go.
	var x [3]int // initialized in the zero value of te class of the array, in this case 0

	var y = [3]int{1, 2, 3}
	fmt.Printf("\n", x, y)

	var sparseArray = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Printf("\n", sparseArray)
	// you can replace the number of the array size, with the `...`
	var guessSize = [...]int{10, 20, 30, 40, 50}
	fmt.Printf("\n", guessSize)

	//you can use != and ==. Arrays are equals if their size and the contain equal values
	sameArr := [3]int{1, 2, 3}
	sameArr2 := [3]int{1, 2, 3}

	if sameArr == sameArr2 {
		fmt.Printf("Arrays are equal.")
	} else {
		fmt.Printf("Arrays are NOT equal.")
	}

	//go does not have good support for matrix
	//but you can simulate multidimensional arrays
	var matrix = [2][5]int{}
	var matrix2 = [5][2]int{}

	fmt.Printf("\n", matrix)
	fmt.Printf("\n", matrix2)

	//Go arrays are read using bracket syntax

	fmt.Println(sameArr2[0])

	//You cannot read or write past the end of an array or use negative index. (compile error)
	//An out-of-bounds read or write in runtime fails with a panic.

	//Built-in function `len()` takes an array and returns its length
	fmt.Println(len(sameArr2))
	fmt.Println(len(matrix))
	fmt.Println(len(matrix2))

	//yopu cannot use a variable to specify a length of an Array
	//because the array size needs to be resolved at compile time.

}
