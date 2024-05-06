package main

import "fmt"

func counter(start int, end int) {
	fmt.Println(start)
	if start < end {
		counter(start+1, end)
	}

}

func main() {
	counter(1, 6)
}
