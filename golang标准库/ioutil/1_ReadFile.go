package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	fileName := "20190103/ioutil/fmt/test1.txt"
	f,err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("hello xmge")
	defer f.Close()

	// ReadFile reads the fmt named by filename and returns the contents.
	// A successful call returns err == nil, not err == EOF. Because ReadFile
	// reads the whole fmt, it does not treat an EOF from Read as an error
	// to be reported.
	b,err := ioutil.ReadFile(fileName)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(b))

}