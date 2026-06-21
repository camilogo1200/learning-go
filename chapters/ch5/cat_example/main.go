package main

import (
	"io"
	"log"
	"os"
	"unsafe"
)

/*
This example demonstrate how to implement the simple cat but in a better idiomatic way, with better practices
*/

const bufferFileSize = 2048

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("filename not included")
	}

	file, closer, err := getFile(args[1])

	if err != nil {
		log.Fatalf("An error occurred open the file. file ={%s}", args[1])
	}

	defer closer()

	bufferFile := make([]byte, bufferFileSize)
	for {
		count, err := file.Read(bufferFile)
		if count == 0 {
			//log.Printf("count = {%d}, bufferFile = {%s}, EOF.", count, string(bufferFile))
			//log.Printf("count = {%d}, bufferFile = {%s}, EOF.", count, string(bufferFile))
			//zero copy  conversion high performance byte[] to string
			str := unsafe.String(&bufferFile[0], len(bufferFile))
			log.Printf("count = {%d}, bufferFile = {%s}, EOF.", count, str)
			break
		}
		if err != nil {
			if err != io.EOF {
				log.Fatalf("An Error occurred reading the file. file ={%s}", args[1])
			}
			break
		}
		_, err = os.Stdout.Write(bufferFile[:count])
		if err != nil {
			log.Fatalf("An Error occurred writing to Stdout. file ={%s}", args[1])
		}
	}
}

func getFile(filename string) (*os.File, func(), error) {

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

	return file, closer, err
}
