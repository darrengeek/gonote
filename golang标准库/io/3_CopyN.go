package main

import (
	"fmt"
	"io"
	"log"
	"os"
)


func main() {

	f1,err := os.Create("20190102/io/fmt/test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	f2,err := os.Create("20190102/io/fmt/test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	_,err = io.WriteString(f1,"hello xmge")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0,os.SEEK_SET)

	// CopyN copies n bytes (or until an error) from src to dst.
	// It returns the number of bytes copied and the earliest
	// error encountered while copying.
	// On return, written == n if and only if err == nil.
	//
	// If dst implements the ReaderFrom interface,
	// the copy is implemented using it.
	fmt.Println(io.CopyN(f2,f1,10))
}
