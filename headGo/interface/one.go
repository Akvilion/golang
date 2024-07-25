package main

import "fmt"

type Car1 struct{}

type Truck1 struct{}

type weightInterface interface {
	weigth()
}

func (c Car1) weigth() {
	fmt.Println("car weigth = 2000kg")
}

func (t Truck1) weigth() {
	fmt.Println("truck weigth = 35000kg")
}

func (t Truck1) engine() {
	fmt.Println("truck engine = 6.7L")
}

func main() {
	honda := Car1{}
	man := Truck1{}
	a := weightInterface(honda) // перший раз :=
	a.weigth()
	a = weightInterface(man) // друге присвоєнная =
	truck, ok := a.(Truck1)
	if ok {
		truck.engine()
	}
	// більш запутаний варіант того що вище
	// if truck, ok := a.(Truck1); ok {
	// 	truck.engine()
	// }
}
