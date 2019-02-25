package main

import (
	"os"
	"runtime"
	"strings"
	"log"
	"reflect"
	"fmt"
)

func main() {
	_,path,_,_ := runtime.Caller(0)
   path = path[0:strings.LastIndex(path,"/") + 1]
	os.Chdir(path)
	f,err := os.Create("test7.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	fInfo,err  := os.Stat(f.Name())
	if err != nil {
		log.Println(err)
	}

	v :=reflect.ValueOf(fInfo)

	for i:=0;i<v.NumMethod();i++ {
		fmt.Println(v.Method(i))
		v.Method(i).Call(nil)
	}
}
