package main

// go утвердження типа
// go type assertion

// щоб викликати метод типу який не передбачений в інтерфейти
// потрібно використати
// var myVariable objType = interfaceVar.(objType)
// myVariable.methodMissingInUnterface()

import "fmt"

type typeOne string

type typeTwo string

func (t typeOne) nnn() {
	fmt.Println("type 1 nnn")
}

func (t typeOne) fff() {
	fmt.Println("type 1 fff")
}

func (t typeTwo) nnn() {
	fmt.Println("type 2 nnn")
}

type RRR interface {
	nnn()
}

func main() {
	var a RRR = typeOne("alan")
	var b typeOne = a.(typeOne)
	b.fff()
}
