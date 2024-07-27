package main

import "fmt"

// до го 1.18 використовувати пустий інтерфейс interface{}
// а з го 1.18 використовують any

func getKeys(m any) ([]any, error) {
	switch t := m.(type) {
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
