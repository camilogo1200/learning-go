package main

import "fmt"

func main() {

	var mySlice []int

	fmt.Printf("mySlice { values = [%v] - len = [%v] - capacity = [%v]\n", mySlice, len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 0)
	for i := 1; i < 2000000; i *= 2 {
		fmt.Printf("Adding double of the elements = [%v]\n", i)
		for e := 0; e < i; e++ {
			mySlice = append(mySlice, 1)
		}

		lenRatio := 100 - float32(float32(len(mySlice))*100)/float32(cap(mySlice))
		ratio := float32(len(mySlice)) / float32(cap(mySlice))
		fmt.Printf("mySlice { len = [%v] - capacity = [%v] - diff = [%.2f] - ratio =[%.2f]\n ", len(mySlice), cap(mySlice), lenRatio, ratio)
	}

	//a slice grow rate is defined by the equation (current_capacity + 768) / 4 when the items are more than 256 -? converges at 25%
	//a slice grow double its capacity when the size is less than 256

}
