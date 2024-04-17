package magazine

import "fmt"

type Subscriber struct {
	name   string
	Rate   float64
	active bool
}

func printInfo(s *Subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.name = name
	s.Rate = 5.99
	s.active = true
	return &s
}
func applyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

// func main() {
// 	subscriber1 := defaultSubscriber("Aman Singh")
// 	applyDiscount(&subscriber1)
// 	printInfo(subscriber1)
// 	subscriber2 := defaultSubscriber("Beth Ryan")
// 	printInfo(subscriber2)
// }
