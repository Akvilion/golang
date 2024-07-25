package main

import "fmt"

func one() {
	// asciiString := "ABCDE"
	utf8String := "БГДЖИ"
	for _, v := range utf8String {
		fmt.Println(string(v))
	}
}

func main() {
	one()
}
