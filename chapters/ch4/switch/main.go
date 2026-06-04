package main

import "fmt"

func main() {

	//Go has a switch statement
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "Is a long word!")
		}
	}

	//you can break the switch using break or break with label
	//example of expresions && break
	for i := 0; i < 10; i++ {
		switch {
		case i == 9:
			fmt.Println(i, "Exit the case expression! not the loop!")
			break
		case i%2 == 0:
			fmt.Printf("i {%d} is even!\n", i)
		case i%3 == 0 && i%2 != 0:
			fmt.Printf("i {%d} is divisible by 3 but not by 2!\n", i)

		default:
			fmt.Printf("i {%d} is boring.\n", i)
		}
	}

	//exit the switch and the loop using labels

loop:
	for i := 0; i < 10; i++ {
		switch {
		case i == 6:
			fmt.Println(i, "Exit the loop!")
			break loop
		case i%2 == 0:
			fmt.Printf("i {%d} is even!\n", i)
		case i%3 == 0 && i%2 != 0:
			fmt.Printf("i {%d} is divisible by 3 but not by 2!\n", i)

		default:
			fmt.Printf("i {%d} is boring.\n", i)
		}
	}
}
