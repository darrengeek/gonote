package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f1,_ := os.Create("20190104/bufio/file/test5_1.txt")
	f2,_ := os.Create("20190104/bufio/file/test5_2.txt")

	f1.WriteString("xmge 11")
	f1.Seek(0,io.SeekStart)

	r := bufio.NewReader(f1)
	w := bufio.NewWriter(f2)
	rw := bufio.NewReadWriter(r,w) //相当于io.TeeReader


	// ReadLine is a low-level line-reading primitive. Most callers should use
	// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
	//
	// ReadLine tries to return a single line, not including the end-of-line bytes.
	// If the line was too long for the buffer then isPrefix is set and the
	// beginning of the line is returned. The rest of the line will be returned
	// from future calls. isPrefix will be false when returning the last fragment
	// of the line. The returned buffer is only valid until the next call to
	// ReadLine. ReadLine either returns a non-nil line or it returns an error,
	// never both.
	//
	// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
	// No indication or error is given if the input ends without a final line end.
	// Calling UnreadByte after ReadLine will always unread the last byte read
	// (possibly a character belonging to the line end) even if that byte is not
	// part of the line returned by ReadLine.
	b,isProfix,_ := rw.ReadLine()
	fmt.Println(string(b))
	fmt.Println(isProfix)
	rw.WriteString("hello xmge")
	rw.Flush()
}
