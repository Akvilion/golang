package main

import "fmt"

func main() {
	xi := []int{1, 2, 3}
	fmt.Println(xi)
	changeSlice(xi, 1)
	fmt.Println(xi)
}

func changeSlice(a []int, b int) {
	a[b] *= 2
}
