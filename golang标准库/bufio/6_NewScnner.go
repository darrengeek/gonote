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

	input := "foo  bar   baz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
