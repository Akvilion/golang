package main

import (
	"fmt"
)

func main() {
	// log.Fatal("something") // завершить програму після виводу
	// fmt.Println("hello")   // це не буде виведено

	err := fmt.Errorf("123")
	fmt.Println(err)

}
