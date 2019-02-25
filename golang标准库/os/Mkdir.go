package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.MkdirAll("20190105/os/file/test2",0777))
}
