package main

import (
	"fmt"
	"xxx/src/github.com/headfirstgo/datafile"
)

func main() {
	res, _ := datafile.GetStrings("../datafile/votes.txt")
	fmt.Println(res)
}
