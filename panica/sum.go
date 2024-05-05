package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Open file", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Close File")
	file.Close()
}
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		number, err := strconv.ParseFloat(scaner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	CloseFile(file)
	if scaner.Err() != nil {
		return nil, scaner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
