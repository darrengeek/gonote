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

	f,err := os.Create("20190104/bufio/fmt/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(f,"hello xmge")
	f.Seek(0,io.SeekStart)

	// NewReader returns a new Reader whose buffer has the default size.
	// defaultBufSize = 4096
	r := bufio.NewReader(f)
	c,err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(c))
}
