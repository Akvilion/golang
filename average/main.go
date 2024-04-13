package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var a float64
	for _, value := range numbers {
		a += value
	}
	fmt.Println(a / 3)
}
