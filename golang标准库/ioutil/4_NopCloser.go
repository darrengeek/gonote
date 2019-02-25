package main

import (
	"io/ioutil"
	"os"
	"io"
	"fmt"
)

func main() {
	f,_ := os.Create("20190103/ioutil/fmt/test3")
	io.WriteString(f,"hello xmge")
	f.Seek(0,io.SeekStart)

	// NopCloser returns a ReadCloser with a no-op Close method wrapping
	// the provided Reader r.
	r := ioutil.NopCloser(f)
	buffer := make([]byte,24)
	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))
}