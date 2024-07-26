package main

import "fmt"

type moveInt interface {
	move()
}

type Car struct {
	mover moveInt
}

func (c Car) Run() {
	c.mover.move()
}

type engine struct{}

func (e engine) move() {
	fmt.Println("Move")
}

func main() {
	eng := engine{}
	x := Car{mover: eng}
	x.Run()
}
