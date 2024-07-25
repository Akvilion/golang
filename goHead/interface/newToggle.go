package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	f := Switch("on")     // ініціалізуємо Switch := потрібно замість var f Switch // Switch("on")
	var a Toggleable = &f // берез вказівник на switch
	a.toggle()
	a.toggle()
}
