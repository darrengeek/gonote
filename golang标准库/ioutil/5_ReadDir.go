package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	// ReadDir reads the directory named by dirname and returns
	// a list of directory entries sorted by filename.
	fInfos,err  := ioutil.ReadDir("20190103/ioutil")
	if err != nil {
		log.Fatal(err)
	}

	// fileInfo 如何获取
	for _,info := range fInfos {
		fmt.Println("=============")
		fmt.Printf("info.Name():%s\n",info.Name())
		fmt.Printf("info.IsDir():%t\n",info.IsDir())
		fmt.Printf("info.Mode():%s\n",info.Mode())
		fmt.Printf("info.ModeTime():%s\n",info.ModTime())
		fmt.Printf("info.Size():%d\n",info.Size())
		fmt.Printf("info.Sys():%s\n",info.Sys())
		fmt.Println("=============")
	}
}