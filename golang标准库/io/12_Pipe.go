package main

import (
	"fmt"
	"io"
	"time"
)

func main() {


	// Read implements the standard Read interface:
	// it reads data from the pipe, blocking until a writer
	// arrives or the write end is closed.
	// If the write end is closed with an error, that error is
	// returned as err; otherwise err is EOF.


	// 相当于一个无缓存通道
	r,w := io.Pipe()
	go func() {
		for {
			buffer := make([]byte,10)
			fmt.Println(r.Read(buffer))
	}
	}()

	go func() {
		for i := 0;i <= 100;i++{
			fmt.Println("w 写入数据",i)
			w.Write([]byte(fmt.Sprintf("%d",i)))
		}
	}()


	time.Sleep(10 * time.Second)
}
