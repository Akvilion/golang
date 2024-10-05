package main

import "fmt"

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

type animal interface{}

func main() {
	d1 := dog{"dog1", 1}
	d2 := dog{"dog2", 2}
	c1 := cat{"cat3", 3}

	xd := []animal{d1, d2, c1}
	infoCreature(xd)
}

type createure interface {
	animal
}

func info[T createure](t T) string {
	return fmt.Sprintf("%v", t)
}

func infoCreature[T createure](tt []T) {
	for _, v := range tt {
		fmt.Println(info(v))
	}
}
