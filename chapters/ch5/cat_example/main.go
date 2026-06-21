package main

import (
	"log"
	"os"
)

/*
This example demonstrate how to implement the simple cat but in a better idiomatic way, with better practices
*/
func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("filename not included")
	}

	file, closer, err := getFile(args[1])

}

func getFile(filename string) (os.File, func(), error) {

	if filename == "" {
		log.Fatalln("invalid filename")
	}

	file, err := os.Open(filename)

	closer := func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("There was an error closing the file [%s]. error = {%v}", filename, err)
			return
		}
	}
}
