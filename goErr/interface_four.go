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
	fmt.Println("Engine move")
}

type electricEngine struct{}

func (e electricEngine) move() {
	fmt.Println("Electric engine move")
}

func main() {
	eng := engine{}
	eEng := electricEngine{}

	car1 := Car{mover: eng}
	car2 := Car{mover: eEng}

	car1.Run() // Output: Engine move
	car2.Run() // Output: Electric engine move
}

// У цьому прикладі ви можете легко замінити engine на electricEngine без зміни коду в методі Run класу Car.
// Це ілюструє гнучкість, яку забезпечують інтерфейси.
// Таким чином, перевага програмування на рівні інтерфейсів полягає в можливості змінювати реалізації інтерфейсу
// без зміни клієнтського коду, що робить код більш гнучким, тестованим і модульним.
