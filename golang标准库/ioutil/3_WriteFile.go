package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	fmt.Println(ioutil.WriteFile("20190103/ioutil/fmt/test3.txt",[]byte("hello xmge"),0666))
}
