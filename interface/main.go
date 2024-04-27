package main

import "xxx/src/mypkg"

func main() {
	// var a mypkg.MyInterface // оголошуємо тип змінної
	a := mypkg.MyType(5) // MyType підтримує MyInterface
	a.MethodWithoutParameters()

}
