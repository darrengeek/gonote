package main

import (
	"fmt"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Human struct{}

func (h *Human) ShowA() {
	fmt.Println("Human showA")
}

type Teacher struct {
	Human
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}