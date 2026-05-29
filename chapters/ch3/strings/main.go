package main

import "fmt"

func main() {

	var x rune = '🔥'
	var heartBlue rune = '💙'
	var ice rune = '🧊'

	fmt.Printf(" fire = [%c %c %c ]\n", x, heartBlue, ice)
	fmt.Printf(" fire = [%q %q %q ] \n", x, heartBlue, ice)
	fmt.Printf(" fire = [%d %d %d ]\n", x, heartBlue, ice)

	sentence := "Hello, 🫐🌽🥑🫛🥦🍆, there 🚀 !"
	runes := []rune(sentence)
	bytes := []byte(sentence)

	fmt.Println(sentence)

	for i := 0; i < len(sentence); i++ {
		fmt.Printf("[%c]", sentence[i])
	}
	fmt.Printf("\n")
	for _, value := range sentence {
		fmt.Printf("[%c]", value)
	}
	fmt.Printf("\n")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("[%c]", runes[i])
	}
	fmt.Printf("\n")
	for _, value := range runes {
		fmt.Printf("[%c]", value)
	}
	fmt.Printf("\n")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("[%c]", bytes[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("[%q]", bytes[i])
	}
	fmt.Printf("\n")
	fmt.Println(string(bytes))
	fmt.Println(string(runes))
}
