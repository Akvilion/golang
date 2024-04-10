package main

import "fmt"

func one(word string) {
	fmt.Println(word)
}

func three(one int, two int) int {
	return one * two
}

func two(text string, counter int) {
	for i := 0; i < counter; i++ {
		fmt.Println(text)
	}
}

func main() {
	// one("hello world")
	// two("alan", 4)
	a := three(2, 4)
	fmt.Println(a)
}
