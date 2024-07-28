package main

import "fmt"

func ine[K comparable, V any](x map[K]V) ([]K, []V) {
	var keys []K
	var values []V
	for k, v := range x {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func main() {
	myMap := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 40,
	}

	keys, values := ine(myMap)
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)
}
