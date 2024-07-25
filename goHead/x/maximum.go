package main

import "fmt"

func maximum(args ...int) int {
	max := 0
	for _, value := range args {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	a := maximum(4, 2, 3, 1)
	fmt.Println(a)
}
