package main

import "fmt"

type car struct {
	model  string
	engine string
	fuel   bool
}

func printStr(s *car) {
	fmt.Println(s.model)
	fmt.Println(s.engine)
	fmt.Println(s.fuel)
}

func populateCar(model string) *car {
	var s car
	s.model = model
	s.engine = "electric"
	s.fuel = false
	return &s
}

func changeMotor(c *car) {
	c.engine = "engineV2"
}

func main() {
	model := populateCar("tesla")
	changeMotor(model)
	printStr(model)
}
