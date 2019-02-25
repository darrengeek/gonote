package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f,err := os.Create("20190104/bufio/fmt/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(f,"hello xmge")
	f.Seek(0,io.SeekStart)

	// NewReaderSize returns a new Reader whose buffer has at least the specified
	// size. If the argument io.Reader is already a Reader with large enough
	// size, it returns the underlying Reader.
	r := bufio.NewReaderSize(f,5)
	c,err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(c))
}
