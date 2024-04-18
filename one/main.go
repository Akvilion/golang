package main

import (
	"fmt"
	"xxx/src/github.com/headfirstgo/magazine"
)

func main() {
	// var s magazine.Subscriber
	// s.rate = 1.99
	// fmt.Println(s.rate)

	var s magazine.Subscriber
	s.Rate = 1.99
	fmt.Println(s)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000

	var address = magazine.Address{Street: "Stusa", City: "Dashava", State: "Lviv", PostalCode: "430001"}
	employee.HomeAddress = address

	fmt.Println(employee)

}
