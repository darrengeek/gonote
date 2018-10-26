package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
}

func (this Dog)Speak()  {
	fmt.Println("汪 汪 汪")
}

type Cat struct {
}

func (this Cat)Speak()  {
	fmt.Println("喵 喵 喵")
}


func main() {
	var zoo = []Animal{}
	zoo = append(zoo, new(Dog),new(Cat))
	for _,a := range zoo {
		a.Speak()
	}
}
