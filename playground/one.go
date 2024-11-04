package main

import "fmt"

func main() {
	ch := make(chan int)

	go one(ch)

	<-ch

	fmt.Println(2)

}

func one(ch chan int) {
	defer close(ch)
	fmt.Printf("one -> ch")
}
