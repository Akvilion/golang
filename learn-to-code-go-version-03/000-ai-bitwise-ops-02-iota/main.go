package main

import "fmt"

const (
	a int = iota // починає з 0 і для кожного значення додає 1
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
	fmt.Println(d) // 3
	fmt.Println(e) // 4
	fmt.Println(f) // 5
}
