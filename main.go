package main

import "fmt"

func main() {

	var v interface{} = 1.0
	v1, ok := v.(float32)
	fmt.Println(v1)
	fmt.Println(ok)
}
