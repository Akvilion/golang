package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/GoesToEleven")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer resp.Body.Close() // Закриваємо тіло відповіді після використання

	type G struct {
		Name        string `json:"name"`
		PublicRepos int    `json:"public_repos"` // Додали правильний тег json
	}

	var f G
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&f)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Name: %s\nPublic Repos: %d\n", f.Name, f.PublicRepos)
}
