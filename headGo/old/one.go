package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a grade")
	data, err := reader.ReadString('\n')
	if err != nil {
		// fmt.Println("some error")
		log.Fatal(err)
	}
	data = strings.TrimSpace(data) // removes leading and trailing white space from a string

	grade, err := strconv.ParseFloat(data, 64)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(grade)

	// fmt.Println("you value = ", grade)

	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failed"
	}

	fmt.Println("A grade of", grade, "is", status)

}
