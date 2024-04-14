package datafile

// package main

import (
	"bufio"
	"os"
	"strconv"
)

// читає значення float64 з кожної строки файла
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	//i := 0
	scanner := bufio.NewScanner(file) // читаємо файл в scanner
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64) // перетворення кожної text в float
		if err != nil {
			return numbers, err
		}
		// i++
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

// func main() {
// 	num, _ := GetFloats("data.txt")
// 	fmt.Println(num)
// }
