package main

import "fmt"

type Date struct {
	Day    int
	Mounth int
	Year   int
}

func (d *Date) SetYear(year int) { // метод отримує копію (d Date) -> (d *Date)
	d.Year = year
}

func main() {
	date := Date{}
	date.SetYear(1991)
	fmt.Println(date.Year) // 0 - виведе
}
