package main

import "fmt"

func main() {
	// is much more efficient to size from the beginning the slice for that matter you use `make`
	/*
		make allows you to specify the type, length and optionally the capacity.
	*/

	mySlice := make([]int, 5)
	fmt.Printf("mySlice { values = [%v] - len = [%v] - capacity = [%v]\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 10)
	fmt.Printf("mySlice { values = [%v] - len = [%v] - capacity = [%v]\n", mySlice, len(mySlice), cap(mySlice))
	myInitialCapacitySlice := make([]int, 5, 10)
	fmt.Printf("myInitialCapacitySlice { values = [%v] - len = [%v] - capacity = [%v]\n", myInitialCapacitySlice, len(myInitialCapacitySlice), cap(myInitialCapacitySlice))
	//zero length but a capacity greater than zero
	sliceLenZero := make([]int, 0, 100)
	fmt.Printf("sliceLenZero { values = [%v] - len = [%v] - capacity = [%v]\n", sliceLenZero, len(sliceLenZero), cap(sliceLenZero))

	//emptying a slice
	s := []string{"first", "second", "third"}
	fmt.Printf("s { values = [%v] - len = [%v] - capacity = [%v]\n", s, len(s), cap(s))
	clear(s)
	fmt.Printf("s { values = [%v] - len = [%v] - capacity = [%v]\n", s, len(s), cap(s))
	s = append(s, "four")
	fmt.Printf("s { values = [%v] - len = [%v] - capacity = [%v]\n", s, len(s), cap(s))
	s[0] = "first"
	fmt.Printf("s { values = [%v] - len = [%v] - capacity = [%v]\n", s, len(s), cap(s))

}
