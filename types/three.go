package main

import "fmt"

type Liters float64
type Mililiters float64
type Gallons float64

func (x Gallons) toLiters() {
	fmt.Println(x * 0.286)
}

func main() {
	xxx := Gallons(11.7)
	xxx.toLiters()

}
