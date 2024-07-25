package main

import "fmt"

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d *Date) SetYear(year int) { // метод отримує копію (d Date) -> (d *Date)
	d.Year = year
}

func (d *Date) SetMonth(month int) { // метод отримує копію (d Date) -> (d *Date)
	d.Month = month
}

func (d *Date) SetDay(day int) { // метод отримує копію (d Date) -> (d *Date)
	d.Day = day
}

func main() {
	// date := Date{}
	// date.SetYear(1991)
	// fmt.Println(date.Year) // 0 - виведе

	date := Date{}
	date.SetYear(2019)
	date.SetMonth(5)
	date.SetDay(27)
	fmt.Println(date)
}
