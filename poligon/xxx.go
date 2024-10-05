package main

import "fmt"

const animalInfo string = `This animal's info is`

type dog struct {
	first string
	age   int
}

type cat struct {
	first string
	age   int
}

type animal interface {
}

func (a animal) voice() {
	fmt.Println(a.first)
}

func main() {
	d1 := dog{"Rover", 7}
	d2 := dog{"Rufus", 8}
	c1 := cat{"Fluffy", 42}
	c2 := cat{"Buffy", 43}

	// using interface
	xa := []animal{d1, d2, c1, c2}

}

// generics

func info[T dog | cat](t T) string {
	return fmt.Sprintf("%v", t)
}

func prtInfo[T dog | cat](tt []T) {
	for i, t := range tt {
		fmt.Printf("%d - %s \n", i, info[T](t))
	}
}
