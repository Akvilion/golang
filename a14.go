package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 2
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(c)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
}
