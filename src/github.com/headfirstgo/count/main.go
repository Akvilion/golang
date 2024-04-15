package main

import (
	"fmt"
	"log"
	"xxx/src/github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt") // "../datafile/votes.txt"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
