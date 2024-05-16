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

func readFile() {
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

func writeFile() { // функція повністю перезаписує файл
	fileName := "aardvark.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("Some text1111111111"))
	check(err)
	err = file.Close()
	check(err)
}

func main() {
	writeFile()
}
