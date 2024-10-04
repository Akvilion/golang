package main

import "fmt"

func main() {
	do(21)
	do("hello")
	do(true)
}

func do(x any) {
	fmt.Printf("Type of x is %T\n", x)
	switch v := x.(type) { // така конструкція x.(type) має зміст тільки в switch
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
