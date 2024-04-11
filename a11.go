package main

import (
	"fmt"
	"myproject/calc"
)

func main() {
	a := calc.Add(1, 2)
	fmt.Println(a)
	b := calc.Subtract(7, 3)
	fmt.Println(b)
}
