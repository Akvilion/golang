package main

import "fmt"

type car struct {
	model  string
	wheels int
}

func one(c car) {
	c.model = "tesla"
	fmt.Print(c)
}

func main() {
	var t car
	t.wheels = 4
	one(t)
}
