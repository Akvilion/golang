package main

import "fmt"

type car struct {
	wheels int
	model  string
}

func showMe(c car) {
	fmt.Println("model", c.model)
	fmt.Println("wheels", c.wheels)
}

func info(name string) car {
	var c1 car
	c1.model = name
	c1.wheels = 4
	return c1
}

func main() {
	c := info("tesla")
	c.wheels = 3
	showMe(c)
}
