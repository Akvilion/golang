package main

import "fmt"

type Foo struct {
	name string
	age  int
}

type Store struct {
	m map[string]*Foo
}

func (s Store) Put(id string, foo *Foo) {
	s.m[id] = foo
}

func main() {
	f := Store{m: make(map[string]*Foo)}
	foo1 := &Foo{name: "one", age: 1}
	foo2 := &Foo{name: "two", age: 2}
	f.Put("first", foo1)
	f.Put("second", foo2)
	for key, foo := range f.m {
		fmt.Printf("Key: %s, Foo Name: %s, Foo Value: %d\n", key, foo.name, foo.age)
	}

}
