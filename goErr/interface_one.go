package main

import "fmt"

type one interface {
	move()
}

type Car struct {
}

type Truck struct {
}

func (t Truck) move() {
	fmt.Println("Truck move")
}

func (t Car) move() {
	fmt.Println("Car move")
}

func main() {
	var car one = Car{}
	var truck one = Truck{}
	car.move()
	truck.move()
}
