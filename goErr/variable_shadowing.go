package main

import (
	"fmt"
)

func one() (int, error) {
	var error error
	if error != nil {
		return 0, error
	}
	return 1, nil
}

func shadowing() {
	x := 2
	if true {
		x, err := one() // в цьому випадку локальна змінна x затіняє глобальну x, тобто глобальна x далі буде рівна 2 а не 1 як нам і треба
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(x)
	}
	fmt.Println(x)
}

func shadowingFixOne() {
	x := 2
	if true {
		temp, err := one()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(temp)
		x = temp // локальна змінна присвоюється в глобальну, викидуємо значення на 1 рівень вище
	}
	fmt.Println(x)
}

func shadowingFixTwo() {
	var x int
	x = 2
	var err error // якщо ми оголошуємо змінну так
	if true {
		x, err = one() // то тут тоді можна використовувати = а не :=
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(x)
	}
	fmt.Println(x)
}

func main() {
	shadowing()
	shadowingFixOne()
	shadowingFixTwo()
}
