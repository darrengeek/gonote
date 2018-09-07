package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Te struct {
}

func (t *Te) ShowA() {
	fmt.Println("TEEEEEEEEEEEEEEE")
}

type Teacher struct {
	People
	Te
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
