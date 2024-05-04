package main

import "fmt"

type MyStr string

func One(a interface{}) {
	fmt.Println(a)
}

func main() {
	//a := MyStr{}
	One("3.14")
	One(123)
	One(true)
	One(MyStr("fff"))
}
