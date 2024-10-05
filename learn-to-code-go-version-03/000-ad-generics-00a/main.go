package main

import "fmt"

func main() {
	loveIt("Hello world")
	loveIt(7)
	loveIt(true)
	loveIt(4.2)
}

// func loveIt[T any](a T) {
// 	fmt.Println(a)
// }

func loveIt[T interface{}](a T) { // ми вказуємо що фуднкція може приймати любий тип через generic а потім передаємо змінну цього типу (a T)
	fmt.Println(a)
}

// [T any] — це generic конструкція, яка дозволяє створювати функції або типи, що можуть працювати з будь-якими типами
