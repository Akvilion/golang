package main

import "fmt"

type myType int

func (m *myType) mul() {
	*m *= 2
}

func main() {
	a := myType(3)
	a.mul()
	fmt.Println(a)
}
