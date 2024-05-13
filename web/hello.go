package main

import (
	"log"
	"net/http"
)

func writeMessage(response http.ResponseWriter, test string) {
	message := []byte(test)
	_, err := response.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(response http.ResponseWriter, request *http.Request) {
	writeMessage(response, "Hello World")
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
