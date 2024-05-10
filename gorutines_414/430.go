package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string, chanel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(len(body))
	chanel <- len(body)
}

func main() {
	myChan := make(chan int)
	go responseSize("https://example.com/", myChan)
	go responseSize("https://golang.org/", myChan)
	go responseSize("https://golang.org/doc", myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
}
