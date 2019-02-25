package main

import (
	"os"
	"log"
	"io"
)

func main() {


	// Create creates the named file with mode 0666 (before umask), truncating
	// it if it already exists. If successful, methods on the returned
	// File can be used for I/O; the associated file descriptor has mode
	// O_RDWR.
	// If there is an error, it will be of type *PathError.
	//	func Create(name string) (*File, error) {
	//	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	//	}

	err := os.Mkdir("20190105/os/file/",0666)
	if err != nil {
		log.Fatal(err)
	}


	f,err := os.Create("20190105/os/file/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("hello xmge")
	f.Seek(0,io.SeekStart)
}