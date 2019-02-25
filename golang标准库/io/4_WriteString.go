package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// Create : OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// Open   : OpenFile(name, O_RDONLY, 0)
	f1,err := os.Create("20190102/io/fmt/test4.txt")
	defer f1.Close()
	if err != nil {
		log.Fatal(err)
	}

	// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
	// If w implements a WriteString method, it is invoked directly.
	// Otherwise, w.Write is called exactly once.
	// 读写文件会更改 文件的游标
	_,err = io.WriteString(f1,"hello xmge")
	if err != nil {
		log.Fatal(err)
	}
}
