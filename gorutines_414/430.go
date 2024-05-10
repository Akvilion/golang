package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string, chanel chan Page) {
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
	// chanel <- len(body)
	chanel <- Page{page: url, size: len(body)}
}

type Page struct {
	page string
	size int
}

func main() {
	myChan := make(chan Page)

	urls := [3]string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	for _, v := range urls {
		go responseSize(v, myChan)
	}
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
}
