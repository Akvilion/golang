package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	for _, v := range &a {
		a[2] = 10
		fmt.Println(v)
	}
}
