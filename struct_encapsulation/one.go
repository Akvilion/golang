package main

import "fmt"

type Date struct {
	Day    int
	Mounth int
	Year   int
}

func main() {
	date := Date{Day: 12, Mounth: 06, Year: 1991}
	fmt.Println(date)
}
