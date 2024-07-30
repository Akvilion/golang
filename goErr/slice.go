package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	s2 := s1[1:3] // верхня межа так як і в пітоні не включається
	fmt.Println(s2)
}
