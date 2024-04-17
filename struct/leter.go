package main

import "fmt"

type car struct {
	model string
	wheel int
}

func main() {
	a := car{model: "tesla", wheel: 4}
	fmt.Println(a)
}
