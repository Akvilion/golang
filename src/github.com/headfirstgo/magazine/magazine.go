package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func printInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}
func defaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}
func applyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

type Employee struct {
	Name   string
	Salary float64
}

// func main() {
// 	subscriber1 := defaultSubscriber("Aman Singh")
// 	applyDiscount(&subscriber1)
// 	printInfo(subscriber1)
// 	subscriber2 := defaultSubscriber("Beth Ryan")
// 	printInfo(subscriber2)
// }
