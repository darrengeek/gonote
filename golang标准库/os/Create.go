package main

import (
	"os"
	"time"
)

func main() {
	os.Stdout.WriteString("\r")
	time.Sleep(1 * time.Second)
	os.Stdout.WriteString("bbbbbbb\r")
	time.Sleep(1 * time.Second)
	os.Stdout.WriteString("ccccccc\r")
}