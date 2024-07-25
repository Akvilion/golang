package main

import "fmt"

func inRange(min int, max int, ll ...int) []int {
	var data []int
	for _, value := range ll {
		if value > min && value < max {
			data = append(data, value)
		}
	}
	return data
}

func main() {
	a := inRange(1, 5, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(a)
}
