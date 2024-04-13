package main

import (
	"fmt"
	"log"
	"xxx/src/github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("C:/Projects/golang/src/github.com/headfirstgo/datafile/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Println("Average:", sum/sampleCount)
}
