package main

import "fmt"

// до го 1.18 використовувати пустий інтерфейс interface{}
// а з го 1.18 використовують any

// мінус такого підходу це те що тип знається не на етапі компіляції а на етапі вже виконання програми

func getKeys(m any) ([]any, error) {
	switch t := m.(type) { // тип взнається вже під час виконання програми
	default:
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		// pass
	}
}
