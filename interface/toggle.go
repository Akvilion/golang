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
	s := Switch("on")
	var a Toggleable = s
	a.toggle()
}
