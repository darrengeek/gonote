package main

import (
	"os"
	"log"
)

func main() {
	err := os.Chdir("20190105/os/file/")
	if err != nil {
		log.Fatal(err)
	}

	// MkdirAll creates a directory named path,
	// along with any necessary parents, and returns nil,
	// or else returns an error.
	// The permission bits perm (before umask) are used for all
	// directories that MkdirAll creates.
	// If path is already a directory, MkdirAll does nothing
	// and returns nil.
	log.Println(os.MkdirAll("test6/test6",7777))

}
