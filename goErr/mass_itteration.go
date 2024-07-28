package main

import "fmt"

func one() {
	x := []int{1, 2, 3, 4}

	for i, v := range x {
		fmt.Println(i, v)
	}
	// 0 1
	// 1 2
	// 2 3
	// 3 4
}

func two() {
	x := map[int]string{1: "one", 2: "two"}
	for i, v := range x {
		fmt.Println(i, v)
	}
	// 2 two
	// 1 one
}

func main() {
	two()
}
