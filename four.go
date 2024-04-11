package main

import (
	"fmt"
)

func one(a int, b int) (int, int, error) {
	if a < 0 {
		return 0, 0, fmt.Errorf("a < 0")
	}
	if a == 0 {
		return 0, 0, fmt.Errorf("a == 0")
	}
	return a, b, nil
}

func main() {
	// log.Fatal("something") // завершить програму після виводу
	// fmt.Println("hello")   // це не буде виведено

	// err := fmt.Errorf("123")
	// fmt.Println(err)

	a, b, err := one(-2, 4)
	fmt.Println(a, b)
	fmt.Println(err)

}
