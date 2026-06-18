package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	BufferFileSize = 2048
)

// cat sample program
func main() {
	//in Go teh cleanup code is attached to the function with the `defer` keyword
	fmt.Printf("total args (%v), args = {%v}\n", len(os.Args), os.Args)
	params := os.Args
	if len(params) < 2 {
		log.Fatalln("filename not included in the params")
	}

	f, err := os.Open(params[1])
	if err != nil {
		log.Fatalf("error opening file [%v], error ={%v}", params[1], err)
	}
	//you can use a function method , or closure with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("fatal error on closing the file. err ={%v}", err)
		}
	}(f)

	data := make([]byte, BufferFileSize)
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("fatal error on read. err ={%v}", err)
			}
			break
		}
		_, err = os.Stdout.Write(data[:count])
		if err != nil {
			log.Fatalf("fatal error on write on stdOut. err ={%v}", err)
		}
	}
	//you can defer multiple functions in a Go function.
	//they run in a last-in, first-out in a stack way basically (LIFO) order

	//the code  within defer functions runs after the return statement
	//You can return values in a `defer`, but there is no a way to read those values
}
