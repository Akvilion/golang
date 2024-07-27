package main

import "fmt"

// в цьому випадку функція зможе приймати любий тип -> x interface{}

func one(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Printf("Type is int with value %d\n", t)
	case string:
		fmt.Printf("Type is string with value %s\n", t)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	one(42)
	one("hello")
	one(3.14)
}
