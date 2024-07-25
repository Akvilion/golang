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

func viewEng(response http.ResponseWriter, request *http.Request) {
	writeMessage(response, "Hello World")
}
func viewUrk(response http.ResponseWriter, request *http.Request) {
	writeMessage(response, "Привіт світ")
}

func main() {
	http.HandleFunc("/eng", viewEng)
	http.HandleFunc("/urk", viewUrk)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
