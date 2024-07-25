package main

import "fmt"

func main() {

	asciiString := "ABCDE"
	utf8String := "БГДЖИ"
	// неправильний варіант, тому, що кирилиця записується в 2 байти
	asciiBytes := []byte(asciiString)
	utf8Bytes := []byte(utf8String)
	asciiBytesPartial := asciiBytes[3:]
	utf8BytesPartial := utf8Bytes[3:]
	fmt.Println(string(asciiBytesPartial))
	fmt.Println(string(utf8BytesPartial))

	// правильний варіант для роботи з кирилицею а не ascii
	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)
	asciiRunesPartial := asciiRunes[3:]
	utf8RunesPartial := utf8Runes[3:]
	fmt.Println(string(asciiRunesPartial))
	fmt.Println(string(utf8RunesPartial))

	// отак правильно ітеруватись, треба по рунах
	for position, currentRune := range asciiString {
		fmt.Printf("%d: %s\n", position, string(currentRune))
	}
	for position, currentRune := range utf8String {
		fmt.Printf("%d: %s\n", position, string(currentRune))
	}
}
