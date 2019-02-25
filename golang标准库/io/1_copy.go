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

	_,err = io.WriteString(f1,"你好啊")
	if err != nil {
		log.Fatal(err)
	}

	// Copy copies from src to dst until either EOF is reached
	// on src or an error occurs. It returns the number of bytes
	// copied and the first error encountered while copying, if any.
	//
	// A successful Copy returns err == nil, not err == EOF.
	// Because Copy is defined to read from src until EOF, it does
	// not treat an EOF from Read as an error to be reported.
	//
	// If src implements the WriterTo interface,
	// the copy is implemented by calling src.WriteTo(dst).
	// Otherwise, if dst implements the ReaderFrom interface,
	// the copy is implemented by calling dst.ReadFrom(src).
	fmt.Println(io.Copy(f2,f1))
}
