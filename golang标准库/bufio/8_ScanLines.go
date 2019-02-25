package main

import (
	"bufio"
	"fmt"
	"strings"
)

/*
	本文会更加侧重于讲解 bufio 包中的 Scanner 扫描器模块，它的主要作用是把数据流分割成一个个标记并除去它们之间的空格。
*/

func main() {

	input := "foo\nbar\nbaz"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// ScanLines is a split function for a Scanner that returns each line of
	// text, stripped of any trailing end-of-line marker. The returned line may
	// be empty. The end-of-line marker is one optional carriage return followed
	// by one mandatory newline. In regular expression notation, it is `\r?\n`.
	// The last non-empty line of input will be returned even if it has no
	// newline.
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
