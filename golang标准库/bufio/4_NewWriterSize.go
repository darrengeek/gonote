package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f,err := os.Create("20190104/bufio/fmt/test4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// NewWriter returns a new Writer whose buffer has the default size.
	// defaultBufSize = 4096

	w := bufio.NewWriterSize(f,1024)
	w.WriteString("hello xmge")
	w.Flush()
}