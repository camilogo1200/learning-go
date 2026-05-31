package main

import (
	"fmt"
	"time"
)

func main() {
	type person struct {
		name      string
		lastName  string
		age       int
		birthdate time.Time
	}

	var john person
	doe := person{}

	gmtMinus5 := time.FixedZone("GMT-5", -5*60*60)

	julia := person{
		"Julia", "Johnson", 32, time.Date(1988, time.November, 21, 11, 58, 00, 00, gmtMinus5),
	}

	loc, err := time.LoadLocation("America/Bogota")
	var specificDate time.Time
	if err == nil {
		specificDate = time.Date(1988, time.November, 21, 11, 58, 0, 0, loc)
	}

	camilo := person{
		name:      "camilo",
		lastName:  "Mackenzie",
		age:       24,
		birthdate: specificDate,
	}

	fmt.Println(john)
	fmt.Println(doe)
	fmt.Println(julia)
	fmt.Println(camilo)

	//Anonymous structs

	//You can also declare a variable implements a struct type without previously giving a name to the struct
	//for example

	pet := struct {
		name string
		kind string
	}{
		name: "Kain",
		kind: "Dog",
	}

	fmt.Printf("%v", pet)
	/*
		The Padding and Alignment Caveat
		Modern 64-bit CPUs read memory in 8-byte chunks (words). If a struct's fields don't fit neatly into these words, the Go compiler inserts "padding" (wasted, invisible bytes) to keep the CPU happy.

			Look at this poorly optimized struct:

			Go
			// Junior Struct Layout
			type BadStruct struct {
			IsActive bool    // 1 byte
			// 7 bytes of PADDING added by compiler
			Score    float64 // 8 bytes
			Age      int32   // 4 bytes
			// 4 bytes of PADDING to align the struct to 8 bytes
		}
			// Total Size: 24 Bytes
			By simply reordering the fields from largest to smallest, we optimize the memory footprint:

			Go
			// Architect Struct Layout
			type GoodStruct struct {
			Score    float64 // 8 bytes
			Age      int32   // 4 bytes
			IsActive bool    // 1 byte
			// 3 bytes of PADDING at the end
		}
			// Total Size: 16 Bytes
	*/
}
