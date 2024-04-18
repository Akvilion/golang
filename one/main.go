package main

import (
	"fmt"
	"xxx/src/github.com/headfirstgo/magazine"
)

func main() {
	// var s magazine.Subscriber
	// s.rate = 1.99
	// fmt.Println(s.rate)

	// var s magazine.Subscriber
	// s.Rate = 1.99
	// fmt.Println(s)

	// var employee magazine.Employee
	// employee.Name = "Joy Carr"
	// employee.Salary = 60000

	// var address = magazine.Address{Street: "Stusa", City: "Dashava", State: "Lviv", PostalCode: "430001"}
	// employee.HomeAddress = address
	// fmt.Printf("%#v\n", employee)

	// fmt.Println(employee)

	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress.Street = "123 Oak St"
	subscriber.HomeAddress.City = "Omaha"
	subscriber.HomeAddress.State = "NE"
	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.HomeAddress.Street)
	fmt.Println("City:", subscriber.HomeAddress.City)
	fmt.Println("State:", subscriber.HomeAddress.State)
	fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.HomeAddress.Street = "456 Elm St"
	employee.HomeAddress.City = "Portland"
	employee.HomeAddress.State = "OR"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("Postal Code:", employee.HomeAddress.PostalCode)
}
