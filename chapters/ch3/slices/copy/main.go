package main

import "fmt"

/*
if you need to create a slice that is independent from the original use the built-in functions `copy`
*/
func main() {

	// Full copy
	s := []int{1, 2, 3, 4, 5, 6}

	new := make([]int, len(s))
	num := copy(new, s)
	fmt.Printf("Elements Copied: %d\n", num)
	fmt.Printf("s { values = [%v] - len = [%v] - capacity = [%v]\n", s, len(s), cap(s))
	fmt.Printf("new { values = [%v] - len = [%v] - capacity = [%v]\n", new, len(new), cap(new))

	//partial copy
	newSlice := make([]int, 2)
	numS := copy(newSlice, s)
	fmt.Printf("Elements Copied: %d\n", numS)
	fmt.Printf("newSlice { values = [%v] - len = [%v] - capacity = [%v]\n", newSlice, len(newSlice), cap(newSlice))

	//partial copy slicing
	slice := make([]int, len(s))
	numSlice := copy(slice, s[3:])
	fmt.Printf("Elements Copied: %d\n", numSlice)
	fmt.Printf("newSlice { values = [%v] - len = [%v] - capacity = [%v]\n", slice, len(slice), cap(slice))

}
