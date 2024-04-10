package main

import "fmt"

func one(word string) {
	fmt.Println(word)
}

func two(text string, counter int) {
	for i := 0; i < counter; i++ {
		fmt.Println(text)
	}
}

func main() {
	one("hello world")
	two("alan", 4)
}
