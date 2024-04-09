package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1

	loopCounter := 2

	x := 0

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)

	success := false
	for x < loopCounter {
		fmt.Println("You have", loopCounter-x, "guesses left.")
		x++
		fmt.Println("Enter a guess number: -> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else if guess == target {
			success = true
			fmt.Println("Success")
			break
		}
	}
	// fmt.Println(x)
	if !success {
		fmt.Println("Sorry. You didn’t guess my number. It was:", target)
	}
}
