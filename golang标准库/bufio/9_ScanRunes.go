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

	// ScanRunes is a split function for a Scanner that returns each
	// UTF-8-encoded rune as a token. The sequence of runes returned is
	// equivalent to that from a range loop over the input as a string, which
	// means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
	// Because of the Scan interface, this makes it impossible for the client to
	// distinguish correctly encoded replacement runes from encoding errors.

	// rune is an alias for int32 and is equivalent to int32 in all ways. It is
	// used, by convention, to distinguish character values from integer values.

	//int32的别名，几乎在所有方面等同于int32
	//它用来区分字符值和整数值

	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
