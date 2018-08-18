package main

import "fmt"

func main() {

	m := map[string]interface{}{"hello": "xmge"}
	v1, ok := m["hello"].(string)
	fmt.Printf("v1:%v,ok:%b", v1, ok)
}
