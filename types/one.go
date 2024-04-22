package main

import "fmt"

type Liters float64
type Galons float64

func main() {
	carFuel := Liters(10.0)
	fmt.Println(carFuel)
}
