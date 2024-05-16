package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileName := "aardvark.txt"
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.FileMode(0600))
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}
