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
	// fmt.Println(lines)

	var names []string
	var counts []int

	for _, line := range lines {
		matched := false
		for i, name := range names {
			if line == name {
				counts[i]++
				matched = true
			}
		}
		if !matched {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
