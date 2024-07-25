package main

import "fmt"

type CustomeErr string

func (c CustomeErr) Error() string {
	return string(c)
}

func main() {
	err := CustomeErr("some err")
	fmt.Println(err)
}
