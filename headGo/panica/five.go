package main

import (
	"errors"
	"fmt"
)

// Функція, яка викликає паніку
func panicFunc() {
	defer fmt.Println("Closing refrigerator after panic")
	fmt.Println("Opening refrigerator for panic")
	panic("refrigerator is empty (panic)")
}

// Функція, яка повертає помилку
func errorFunc() error {
	defer fmt.Println("Closing refrigerator after error")
	fmt.Println("Opening refrigerator for error")
	return errors.New("refrigerator is empty (error)")
}

func main() {
	// Обробка паніки
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Виклик функції, яка викликає паніку
	panicFunc()

	// Виклик функції, яка повертає помилку
	err := errorFunc()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
