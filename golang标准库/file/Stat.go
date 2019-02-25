package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {


	fInfo,err := os.Stat("20190101/fmt")
	if err != nil {
		log.Fatal(err)
	}

	v := reflect.ValueOf(&fInfo)

	for i:=0; i < v.NumMethod(); i ++ {
		fmt.Println(v.Method(i).Call(nil))
	}
}

