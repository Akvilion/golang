package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := 2
	f := runtime.Version()
	fmt.Print("hello", a, f)
}
