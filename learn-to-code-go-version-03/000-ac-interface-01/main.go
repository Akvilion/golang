package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	number int
}

func (p person) speak() {
	fmt.Println("I'm a person - ", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a spy - ", sa.first, sa.number)
}

func newPerson(s string) (*person, error) {
	if s == "" {
		return nil, fmt.Errorf("name is blank")
	}
	return &person{s}, nil
}

func newSA(p person, n int) (*secretAgent, error) {
	if p.first == "" {
		return nil, fmt.Errorf("name is blank")
	}
	return &secretAgent{p, n}, nil
}

type human interface { // інтерфейси дозволяють передавати в функції будь-які типи, які реалізують методи, зазначені в інтерфейсі.
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1, _ := newPerson("Moneypenny")
	sa1, _ := newSA(person{"james"}, 7)
	saySomething(p1) //
	saySomething(sa1)
	// Так, будь-яка структура, яка реалізує метод, визначений в інтерфейсі, може бути передана як аргумент в функцію, що приймає цей інтерфейс.
	// Go використовує duck typing для інтерфейсів: якщо тип реалізує всі методи інтерфейсу, то він автоматично вважається реалізатором цього інтерфейсу, навіть якщо явно не вказано, що тип "імплементує" інтерфейс.
}
