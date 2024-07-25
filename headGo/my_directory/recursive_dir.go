package main

import (
	"fmt"
	"log"
	"os"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

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
	defer reportPanic()
	rec("../my_directory")
}
