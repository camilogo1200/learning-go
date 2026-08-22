package main

import "fmt"

// when using iota the best practice is to define a type based on int that will represent all the valid values
type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func main() {
	fmt.Printf("Uncategorized = {%d} \nPersonal = {%d} \nSpam = {%d}\nSocial = {%d} \nAdvertisements = {%d}", Uncategorized, Personal, Spam, Social, Advertisements)
}
