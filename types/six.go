package main

import "fmt"

type Liters float64
type Milliliters float64
type Galons float64

func (l Liters) ToGallons() Galons {
	return Galons(l * 0.264)
}

func (m Milliliters) ToGallons() Galons {
	return Galons(m * 0.000264)
}

func (g Galons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}
func (g Galons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
	milk := Galons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
}

// 2.000 liters equals 0.528 gallons
// 500.000 milliliters equals 0.132 gallons
