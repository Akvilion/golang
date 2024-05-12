package main

import (
	"log"
	"net/http"
)

func viewHandler(response http.ResponseWriter, request *http.Request) {
	message := []byte("Hello world of GO!")
	_, err := response.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
