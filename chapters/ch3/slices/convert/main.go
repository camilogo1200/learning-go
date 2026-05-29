package main

import "fmt"

/*
Converting arrays to slices, as an example when a functions need or requires a slice
*/
func main() {
	//If you have an array you can take a slice from it using a slice expression.
	//To convert an entire array into a slice, use the `[:]` notation

	xArray := [4]int{1, 2, 3, 4}
	nArray := xArray[:]
	//however they share memory but in this case slice cannot grow due to the origin ( array) of the memory allocation

	fmt.Printf("Address of xArray: %p\n", &xArray)
	fmt.Printf("Address of nArray: %p\n", &nArray[0]) //same address
	fmt.Printf("xArray { values = [%v] - len = [%v] - capacity = [%v]\n", xArray, len(xArray), cap(xArray))
	fmt.Printf("nArray { values = [%v] - len = [%v] - capacity = [%v]\n", nArray, len(nArray), cap(nArray))
	nArray[3] = 300
	nArray = append(nArray, 5, 6)
	fmt.Printf("Address of xArray: %p\n", &xArray)
	fmt.Printf("Address of nArray: %p\n", &nArray[0]) //after append, as the previous was array, this sentence generate a new memory space

	fmt.Printf("xArray { values = [%v] - len = [%v] - capacity = [%v]\n", xArray, len(xArray), cap(xArray))
	fmt.Printf("nArray { values = [%v] - len = [%v] - capacity = [%v]\n", nArray, len(nArray), cap(nArray))

	nArray[0] = 100
	fmt.Printf("xArray { values = [%v] - len = [%v] - capacity = [%v]\n", xArray, len(xArray), cap(xArray))
	fmt.Printf("nArray { values = [%v] - len = [%v] - capacity = [%v]\n", nArray, len(nArray), cap(nArray))

	fmt.Println("\nSlice to Array\n")
	//Convert slices to arrays
	// in this conversion new mem mory is created so changes on the original slice does not affect the new array
	xSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newArray := [9]int(xSlice)
	smallArray := [2]int(xSlice)
	fmt.Printf("xArray { values = [%v] - len = [%v] - capacity = [%v]\n", xSlice, len(xSlice), cap(xSlice))
	fmt.Printf("newArray { values = [%v] - len = [%v] - capacity = [%v]\n", newArray, len(newArray), cap(newArray))
	fmt.Printf("smallArray { values = [%v] - len = [%v] - capacity = [%v]\n", smallArray, len(smallArray), cap(smallArray))

	xSlice[0] = 5000
	fmt.Printf("xArray { values = [%v] - len = [%v] - capacity = [%v]\n", xSlice, len(xSlice), cap(xSlice))
	fmt.Printf("newArray { values = [%v] - len = [%v] - capacity = [%v]\n", newArray, len(newArray), cap(newArray))
	fmt.Printf("smallArray { values = [%v] - len = [%v] - capacity = [%v]\n", smallArray, len(smallArray), cap(smallArray))

}
