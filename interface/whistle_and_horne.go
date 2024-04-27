package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!", w)
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!", h)
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker // без оцеї лінійки нічого не заводиться
	toy = Horn("xxxx")
	toy.MakeSound()
	toy = Whistle("wwww")
	toy.MakeSound()

	play(Horn("aaaa"))
}
