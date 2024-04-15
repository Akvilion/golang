package main

import (
	"fmt"
	"log"
	"xxx/src/github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("../datafile/votes.txt") // "../datafile/votes.txt"
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
