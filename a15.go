package main

import "fmt"

func main() {
	a := [4]int{11, 22, 33, 44}
	fmt.Println(a)
	for i := 0; i < 4; i++ {
		fmt.Println(i, a[i])
	}
}
