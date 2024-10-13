package main

import "testing"

func TestAdder(t *testing.T) {
	want := 4
	got := adder(1, 3)
	if want != got {
		t.Errorf("got %v wanted %v", got, want)
	}
}
