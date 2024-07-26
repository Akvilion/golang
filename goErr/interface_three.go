package main

import "fmt"

type vehicleMove interface {
	move()
}

type Engine struct{}

func (e Engine) move() {
	fmt.Println("Engine moves the car")
}

type Car struct {
	mover vehicleMove
}

func (c Car) Move() {
	c.mover.move()
}

func main() {
	engine := Engine{}
	car := Car{mover: engine}
	car.Move()
}
