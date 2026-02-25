package main

import "fmt"

// declaring immutable variables is done with the const keyword
const x int64 = 64
const (
	idKey   = "id"
	nameKey = " name"
)
const z = 20 * 30

func main() {

	const y = " hello : "
	fmt.Printf("", x, idKey, nameKey, z, y)

	// z = z + 1 -> error cannot assign
}
