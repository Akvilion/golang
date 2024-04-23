package main

import (
	"fmt"
	"log"
	"xxx/src/github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetDay(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(6)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1991)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

}
