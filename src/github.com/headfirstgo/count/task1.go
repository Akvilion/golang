package main

import "fmt"

func main() {
	ranks := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	for key, value := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", key, value)
	}
}
