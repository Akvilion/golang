package main

import "fmt"

func hello(x chan string) {
	x <- "hi"
}

func main() {
	x := make(chan string)
	go hello(x) // має бути горутина щоб дістати chanel
	fmt.Println(<-x)
}
