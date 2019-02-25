package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

func main() {

	f1,err := os.Create("20190102/io/fmt/test10_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	f1.WriteString("1")
	f1.Seek(0,io.SeekStart)

	f2,err := os.Create("20190102/io/fmt/test10_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	f2.WriteString("2")
	f2.Seek(0,io.SeekStart)

	f3,err := os.Create("20190102/io/fmt/test10_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	f3.WriteString("3")
	f3.Seek(0,io.SeekStart)


	// MultiReader returns a Reader that's the logical concatenation of
	// the provided input readers. They're read sequentially. Once all
	// inputs have returned EOF, Read will return EOF.  If any of the readers
	// return a non-nil, non-EOF error, Read will return that error.
	r := io.MultiReader(f1,f2,f3)
	buffer := make([]byte,2)

	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

	fmt.Println(r.Read(buffer))
}