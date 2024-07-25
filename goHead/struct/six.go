package main

import "fmt"

type car struct {
	engine string
}

func setEngine(c *car) {
	c.engine = "v6"
}

func main() {
	var s car
	// щоб поміняти наш s потрібно його передавати через вказівник
	setEngine(&s)
	fmt.Println(s.engine)
}
