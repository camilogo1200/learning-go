package main

import "fmt"

func main() {
	/*
		Slicing Slices
		A slice expression to create a slice form a slice.
		Written inside brackets and consist of a starting offset and ending offset, separated by a color (:)
		if you leave off the start offset 0 is the default
		if you leave off the end offset of the slice, the end of the slice is substituted
		# IMPORTANT:
		- When you take a slice from a slice you are not making a copy of the data. You are sharing memory in two variables
	*/

	mySlice := []int{0, 1, 2, 3, 4}
	myNewSlice := mySlice[2:]
	myNewSlice2 := mySlice[:2]

	fmt.Printf("mySlice { values = [%v] - len = [%v] - capacity = [%v]\n", mySlice, len(mySlice), cap(mySlice))
	fmt.Printf("myNewSlice { values = [%v] - len = [%v] - capacity = [%v]\n", myNewSlice, len(myNewSlice), cap(myNewSlice))
	fmt.Printf("myNewSlice2 { values = [%v] - len = [%v] - capacity = [%v]\n", myNewSlice2, len(myNewSlice2), cap(myNewSlice2))
	fmt.Println("Remember slicing is a way to share memory, NOT copy the values or creating a new structure")
	//check how changing elements on myNewSlice, myNewSlice2 affects mySlice (the original slice)

	myNewSlice[0] = 500
	myNewSlice2[0] = 100
	fmt.Printf("mySlice { values = [%v] - len = [%v] - capacity = [%v]\n", mySlice, len(mySlice), cap(mySlice))

	fmt.Println("append has its own logic - check the append didnt add instead replaces! be careful")
	myNewSlice2 = append(myNewSlice2, 999)
	fmt.Printf("mySlice { values = [%v] - len = [%v] - capacity = [%v]\n", mySlice, len(mySlice), cap(mySlice))

}
