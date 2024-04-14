package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[1:])
	var sum float64 = 0
	arguments := os.Args[1:]
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Println(sum / float64(len(arguments)))
}
