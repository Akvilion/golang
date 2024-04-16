package main

import "fmt"

type car struct {
	wheels int
	model  string
	vin    int
}

func main() {
	var myCar car
	myCar.wheels = 4
	myCar.model = "tesla"
	myCar.vin = 3434

	fmt.Println(myCar)
}
