package main

import (
	"io"
	"log"
	"os"
)

func main() {

	f1,err := os.Create("20190102/io/fmt/test11_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0,io.SeekStart)

	f2,err := os.Create("20190102/io/fmt/test11_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	f2.Seek(0,io.SeekStart)


	f3,err := os.Create("20190102/io/fmt/test11_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	f3.Seek(0,io.SeekStart)


	// MultiWriter creates a writer that duplicates its writes to all the
	// provided writers, similar to the Unix tee(1) command.
	//
	// Each write is written to each listed writer, one at a time.
	// If a listed writer returns an error, that overall write operation
	// stops and returns the error; it does not continue down the list.
	w := io.MultiWriter(f1,f2,f3)
	str := "123"
	w.Write([]byte(str))
}