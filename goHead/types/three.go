package main

import "fmt"

type Liters float64
type Mililiters float64
type Gallons float64

func (g Gallons) toLiters(mul float64) {
	fmt.Println(float64(g) * mul)
}

func main() {
	xxx := Gallons(22.7)
	xxx.toLiters(0.286)

}
