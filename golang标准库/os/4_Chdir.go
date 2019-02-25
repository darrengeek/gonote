package main

import (
	"os"
	"log"
	"fmt"
)

func main() {

	// Chdir changes the current working directory to the named directory.
	// If there is an error, it will be of type *PathError.
	err := os.Chdir("20190105/os/file")
	if err != nil {
		log.Fatal(err)
	}

	// Create creates the named file with mode 0666 (before umask), truncating
	// it if it already exists. If successful, methods on the returned
	// File can be used for I/O; the associated file descriptor has mode
	// O_RDWR.
	// If there is an error, it will be of type *PathError.
	fmt.Println(os.Create("test3.txt"))

}
