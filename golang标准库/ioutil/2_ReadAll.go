package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"io"
)

func main() {
	fileName := "20190103/ioutil/fmt/test2.txt"
	f,err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("hello xmge")
	f.Seek(0,io.SeekStart)

	// ReadAll reads from r until an error or EOF and returns the data it read.
	// A successful call returns err == nil, not err == EOF. Because ReadAll is
	// defined to read from src until EOF, it does not treat an EOF from Read
	// as an error to be reported.
	b,err := ioutil.ReadAll(f)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(b))

}