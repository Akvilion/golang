//package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("data.txt") // file - це вказівник на відкритий файл
// 	if err != nil {                  // помилка при відкритті файлу
// 		log.Fatal(err)
// 	}
// 	scaner := bufio.NewScanner(file) // читає дані з файла
// 	for scaner.Scan() {
// 		fmt.Println(scaner.Text())
// 	}
// 	err = file.Close() // перевірки чи при закритті файла помилка
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if scaner.Err() != nil { // перевірка на помилку при скануванні
// 		log.Fatal(scaner.Err())
// 	}

// }

package datafile
