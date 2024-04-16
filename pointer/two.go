package main

import "fmt"

type car struct {
	vin int
}

func main() {
	var x car
	x.vin = 333
	var y *car = &x // оголошуємо змінну y з вказівником на структуру
	// fmt.Println(*y.vin)   // так не буде працювати
	fmt.Println(y.vin)    // 333
	fmt.Println((*y).vin) // 333
}
