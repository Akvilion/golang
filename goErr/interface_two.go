package main

import "fmt"

type vehicleMove interface { // задаємо інтерфейс
	move()
}

type Engine struct{} // структура яка буде реалізовувати інтерфейс move()

func (e Engine) move() {
	fmt.Println("Engine moves the car")
}

type Car struct {
	mover vehicleMove // структура яка імплементує даний інтерфейс
}

func (c Car) Move() {
	c.mover.move() // метод структури Car який через інтерфейс mover реалізує метод move()
}

func main() {
	engine := Engine{}
	car := Car{mover: engine} // інтерфейсу mover який має реалізувати метод move() передаємо engine який і реалізує інтерфейс move()
	car.Move()
}
