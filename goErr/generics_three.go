package main

import "fmt"

// тут є [k comparable, v any] тому що мапа складається з key, value
func ine[k comparable, v any](x map[k]v) []k {
	var a []k
	for k := range x {
		a = append(a, k)
	}
	return a
}

func main() {
	myMap := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 40,
	}
	keys := ine(myMap)
	fmt.Println(keys) // Output: [Alice Bob Carol]
}
