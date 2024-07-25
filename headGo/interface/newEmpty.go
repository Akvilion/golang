package main

import "fmt"

type MyInterf interface{}

type One struct{} // пустий інтерфейс

func (o One) oneMethod() {
	fmt.Println("Hellos")
}

func Two(a MyInterf) {
	obj, ok := a.(One)
	if ok {
		obj.oneMethod()
	}
}

func main() {
	x := One{}
	Two(x)
}
