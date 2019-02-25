package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1,err := os.Create("20190102/io/fmt/test6.txt")
	if err != nil {
		log.Fatal(err)
	}

	_,err = io.WriteString(f1,"hello xmge")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0,io.SeekStart)

	// ReadAtLeast reads from r into buf until it has read at least min bytes.
	// It returns the number of bytes copied and an error if fewer bytes were read.
	// The error is EOF only if no bytes were read.
	// If an EOF happens after reading fewer than min bytes,
	// ReadAtLeast returns ErrUnexpectedEOF.
	// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
	// On return, n >= min if and only if err == nil.
	// If r returns an error having read at least min bytes, the error is dropped.
	buffer := make([]byte,24)
	fmt.Println(io.ReadAtLeast(f1,buffer,24))
	fmt.Println(string(buffer))
}
