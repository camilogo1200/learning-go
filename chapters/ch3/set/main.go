package main

import "fmt"

func main() {

	//Using maps as sets
	//remember set does not guarantee the order of elements
	//you can use maps to mimic the functionality

	m := map[int]bool{}
	vals := []int{1, 2, 3, 0, 1, 1, 2, 2, 3, 10}

	for _, value := range vals {
		m[value] = true
	}

	fmt.Printf("vals = [%v] - len = [%d]\n", vals, len(vals))
	fmt.Printf("m = [%v] - len = [%d]\n", m, len(m))

	//use as the set

	if m[3] {
		fmt.Println("3 exists in the set")
	}

	if m[20] {
		fmt.Println("20 exists in the set")
	}

}
