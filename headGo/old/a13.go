package main

import (
	"xxx/greeting" // щоб працював інпорт з папки проекта потрібно зробити go mod init xxx
)

func main() {
	//mypackage.MyFunction()
	greeting.Hello()
	// x, y := keyboard.GetFloat()
	// fmt.Println(x)
	// fmt.Println(y)
}
