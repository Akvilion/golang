package main

import "fmt"

func main() {
	var i int
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println("x")
		fmt.Println(i)
	}()
	i++

	close(ch)
	fmt.Println("y")
	fmt.Println(i)
}
