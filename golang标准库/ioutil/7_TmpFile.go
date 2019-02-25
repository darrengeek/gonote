package main

import (
	"io/ioutil"
	"log"
)

func main() {

	// TempFile creates a new temporary fmt in the directory dir
	// with a name beginning with prefix, opens the fmt for reading
	// and writing, and returns the resulting *os.File.
	// If dir is the empty string, TempFile uses the default directory
	// for temporary files (see os.TempDir).
	// Multiple programs calling TempFile simultaneously
	// will not choose the same fmt. The caller can use f.Name()
	// to find the pathname of the fmt. It is the caller's responsibility
	// to remove the fmt when no longer needed.
	f,err := ioutil.TempFile("20190103/ioutil/fmt/","tmp-")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello xmge TMP")
}