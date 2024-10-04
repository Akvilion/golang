package main

import (
	"fmt"
)

func main() {
	var s1 interface{} = "James"
	var s2 any = "Bond"
	var s3 any = 21

	fmt.Println(s1.(string)) // type assertion
	fmt.Println(s2.(string)) // type assertion
	fmt.Println(s3.(int))    // type assertion

	// якщо тип підтверджується то type assertion просто повертає значення
}
