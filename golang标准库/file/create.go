package main

import (
	"os"
	"log"
)

func main() {
<<<<<<< HEAD

	// Create creates the named file with mode 0666 (before umask), truncating
	// it if it already exists. If successful, methods on the returned
	// File can be used for I/O; the associated file descriptor has mode
	// O_RDWR.
	// If there is an error, it will be of type *PathError.
	f,err := os.Create("20190101/file/create.txt")
=======
	f,err := os.Create("20190101/fmt/create.txt")

>>>>>>> 8588e30ae76975977bbfa24776686a3ff748677e
	defer f.Close()
	if err != nil{
		log.Fatal(err)
	}
	f.WriteString("create fileaaa")
}
