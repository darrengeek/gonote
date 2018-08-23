package main

import "fmt"

// ok断言之通道有值
func main() {

	c := make(chan bool,2)

	<- c
	c <- true
	c <- true

	// 通道未关闭
	b1,ok := <-c
	fmt.Printf("b1:%t,ok:%t \n",b1,ok)
	// b1:true,ok:true


	// 通道已关闭 通道中还有值
	close(c)
	b2,ok := <-c
	fmt.Printf("b2:%t,ok:%t \n",b2,ok)
	// b2:true,ok:true

	// 通道已关闭 通道中无值
	b3,ok := <-c
	fmt.Printf("b3:%t,ok:%t \n",b3,ok)
	// b3:false,ok:false

	a := []byte{}
	a.len()
}
