package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"io"
)

func main() {
	f,err := os.Create("20190105/os/file/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.WriteString("hello xmge"))
	f.Seek(0,io.SeekStart)
	f.Close()



	// Open opens the named file for reading. If successful, methods on
	// the returned file can be used for reading; the associated file
	// descriptor has mode O_RDONLY.
	// If there is an error, it will be of type *PathError.
	f,err = os.Open("20190105/os/file/test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	b,err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

