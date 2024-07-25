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
	chanel <- Page{Page: url, Size: len(body)}
}

type Page struct {
	Page string
	Size int
}

func main() {
	myChan := make(chan Page)

	urls := [3]string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	for _, v := range urls {
		go responseSize(v, myChan)
	}
	for i := 0; i < len(urls); i++ {
		page := <-myChan
		fmt.Printf("%s: %d\n", page.Page, page.Size)
	}

	// fmt.Println(<-myChan)
	// fmt.Println(<-myChan)
	// fmt.Println(<-myChan)
}
