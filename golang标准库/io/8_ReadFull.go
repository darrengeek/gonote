package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1,err := os.Create("20190102/io/fmt/test8.txt")
	if err != nil {
		log.Fatal(err)
	}

	_,err = io.WriteString(f1,"hello xmge")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0,io.SeekStart)

	// ReadFull reads exactly len(buf) bytes from r into buf.
	// It returns the number of bytes copied and an error if fewer bytes were read.
	// The error is EOF only if no bytes were read.
	// If an EOF happens after reading some but not all the bytes,
	// ReadFull returns ErrUnexpectedEOF.
	// On return, n == len(buf) if and only if err == nil.
	// If r returns an error having read at least len(buf) bytes, the error is dropped.
	buffer := make([]byte,24)
	fmt.Println(io.ReadFull(f1,buffer))
	fmt.Println(string(buffer))
}
