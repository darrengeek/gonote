package main

import (
	"os"
	"log"
)


// 打开已经存在的文件，若不存在则报错
func main() {
	f,err := os.Open("open.txt")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("open fmt")
}