package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	r := x(a, func(n int) int { return n * 2 })
	fmt.Println(r)

}

func x(a []int, someName func(int) int) []int {
	var f = make([]int, len(a))
	for i, v := range a {
		f[i] = someName(v)
	}
	return f
}
