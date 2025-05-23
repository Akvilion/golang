package main

import "fmt"

type customers int // кастомний тим що базується на int

func main() {
	fmt.Println(sumInt(42, 43))
	fmt.Println(sumFloat64(42.03, 43.04))
	fmt.Println(sum[int](42, 43))
	fmt.Println(sum[float64](42.03, 43.04))

	var yesterday customers = 84
	var today customers = 168
	fmt.Println(sum(yesterday, today))
}

func sumInt(x, y int) int {
	return x + y
}
func sumFloat64(x, y float64) float64 {
	return x + y
}

type numTypes interface {
	~int | ~float64 // буде приймати любий тип який базується на int або float
}

func sum[T numTypes](x, y T) T {
	return x + y
}
