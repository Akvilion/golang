package main

import (
	"fmt"
	"time"
)

func one(x chan string) {

	time.Sleep(2 * time.Second)
	fmt.Println("111", time.Now())
	x <- "a"

	time.Sleep(3 * time.Second)
	fmt.Println("333", time.Now())
	x <- "b"
}

func main() {
	ch := make(chan string)
	fmt.Println("main1", time.Now())
	go one(ch)
	//time.Sleep(5 * time.Second)
	fmt.Println("main2", time.Now())
	fmt.Println(<-ch, time.Now())
	//time.Sleep(5 * time.Second)
	fmt.Println(<-ch, time.Now())
}

// Yes, if one goroutine is waiting to read a value from a channel
// and a second goroutine must send a value to the channel,
// the first goroutine will block (or "freeze") until it receives a value from the channel.
