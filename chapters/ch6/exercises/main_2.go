package main

import "fmt"

func main() {

	slice := []string{"hello", "my", "dear", "friend", "!"}

	UpdateSlice(slice, "pepe!")
	GrowSlice(slice, "potato!")

	fmt.Println("UpdateSlice = {%v}", slice)
	fmt.Println("GrowSlice = {%v}", slice)
}

func UpdateSlice(a []string, b string) {
	lastPosition := len(a)
	a[lastPosition-1] = b

	fmt.Println("UpdateSlice = {%v}", a)
}

func GrowSlice(a []string, b string) {
	a = append(a, b)
	fmt.Println("GrowSlice = {%v}", a)
}
