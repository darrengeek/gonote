package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1,err := os.Create("20190102/io/fmt/test7.txt")
	if err != nil {
		log.Fatal(err)
	}

	_,err = io.WriteString(f1,"hello xmge")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0,io.SeekStart)

	// NewSectionReader returns a SectionReader that reads from r
	// starting at offset off and stops with EOF after n bytes.
	r := io.NewSectionReader(f1,5,24)

	buffer := make([]byte,1024)
	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))
}
