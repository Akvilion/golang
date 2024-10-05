package main

import "fmt"

func main() {
	loveIt("Hello world")
	loveIt(7)
	loveIt(true)
	loveIt2(4.2) // doesn't work
}

type myConstraint interface {
	string | int | bool // тут ми вказуємо які пити приймаються
}

func loveIt[T myConstraint](a T) {
	fmt.Println(a)
}

func loveIt2[T any](a T) { // тут взагалі любий тип
	fmt.Println(a)
}
