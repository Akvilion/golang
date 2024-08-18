package main

import "fmt"

func xxx(x string) {
	fmt.Printf("I am %s\n", x)
}

func main() {
	var x string
	defer func() {
		xxx(x) // аргументи defer вираховуються зразу тому завжди буде пуста строка
	}()
	x = "alamo"
	fmt.Println(x)
}

// func xxx(x *string) {
// 	fmt.Printf("I am %s\n", *x)
// }

// func main() {
// 	var x string
// 	defer xxx(&x) // аргументи defer вираховуються зразу тому завжди буде пуста строка
// 	x = "alamo"
// 	fmt.Println(x)
// }
