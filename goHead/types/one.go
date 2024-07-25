package main

import "fmt"

type Liters float64
type Galons float64

func toLiters(x Galons) Liters {
	return Liters(x * 0.264)
}
func toGalons(x Liters) Galons {
	return Galons(x * 3.78)
}

func main() {
	carFuel := Liters(10.0)

	// busFuel = galons(22.2)

	carFuel += toLiters(Galons(11.1))
	fmt.Println(carFuel)

}
