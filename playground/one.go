package main

import "fmt"

// можна використовувати як локально (всередині функцій) так поза ними
var f int = 22

// можна використувувати тільки всередині функцій
// f:=22

func main() {
	ch := make(chan int)
	go one(ch)

	<-ch

	fmt.Println(2)

}

func one(ch chan int) {

	fmt.Println(f)
	defer close(ch)
	fmt.Printf("one -> ch")
}
