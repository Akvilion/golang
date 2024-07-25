package main

import "fmt"

func main() {
	var a int = 2

	b := &a
	var b *int = &a
	// ці два записи аналогічні

	fmt.Println(b)
	fmt.Println(*b)
}
