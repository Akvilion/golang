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

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}

// 2.000 liters equals 0.528 gallons
// 500.000 milliliters equals 0.132 gallons
