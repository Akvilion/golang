package main

import "fmt"

// отримуємо тип вказівник
func one(x *int) int {
	return *x // повертаємо значення вказівника
}

func two(x *int) *int {
	*x = 3
	return x
}

func main() {
	a := 2

	r := one(&a)

	fmt.Println(r)

	f := two(&a)
	fmt.Println(f)
	fmt.Println(a)

}
