package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil { // помилка при відкритті файлу
		log.Fatal(err)
	}
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		fmt.Println(scaner.Text())
	}
	err = file.Close() // перевірки чи при закритті файла помилка
	if err != nil {
		log.Fatal(err)
	}
	if scaner.Err() != nil {
		log.Fatal(scaner.Err())
	}

}
