package main

import (
	"fmt"
	"log"
	"os"
)

func rec(dir_name string) {
	files, err := os.ReadDir(dir_name)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory", file.Name())
			rec(file.Name())
		} else {
			fmt.Println("File", file.Name())
		}
	}
}

func main() {
	rec("../my_directory")
}
