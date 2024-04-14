package main

import "fmt"

func one(x ...int) {
	fmt.Println(x)
}

func main() {
	one(1, 2, 3, 4)
}
