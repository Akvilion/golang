package main

import "fmt"

// тут ми міняємо значення а в функції через вказівник

func one(a *int) {
	*a = 3
}

func main() {
	a := 2

	one(&a)

	fmt.Println(a)

}
