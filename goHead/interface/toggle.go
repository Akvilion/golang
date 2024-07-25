package main

// просто переключає з "on" на "off"

import "fmt"

type Switch string

func (s Switch) toggle() {
	if s == "on" {
		s = "off"
	} else {
		s = "on"
	}
	fmt.Println(s)
}

type Toggleable interface {
	toggle()
}

func main() {
	// s := Switch("on") // літерал :=
	// var a Toggleable = s
	// a.toggle()

	var a Toggleable = Switch("on") // працює тому що Switch має метод toggle() який потребує Toggleable інтерфейс
	a.toggle()

	f := Switch("off") // якщо отак записати то "f" не буде мати тип Toggleable інтерфейс!!!
	f.toggle()

}
