package main

import "fmt"

type car struct {
	vin int
}

func main() {
	var x car
	x.vin = 333
	var y *car = &x // оголошуємо змінну з вказівником на структуру (адрес структри менший за розмір всієї структури)
	// fmt.Println(*y.vin)   // так не буде працювати
	y.vin = 444           // вказівник на структуру дає доступ до модифікації оригінальних даних
	fmt.Println(y.vin)    // 333 // це і є аналог (*y).vin
	fmt.Println((*y).vin) // 333
}
