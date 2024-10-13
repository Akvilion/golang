package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(2, 3)
	want := 5
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(2, 3)
	want := 6
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

// test table
func TestSumTable(t *testing.T) {
	tests := []struct { // літеральне оголошення тобто зразу ініціалізується
		xi   []int
		want int
	}{
		{[]int{2, 3}, 5}, // зріз може мати не тільки якийсь один тип int string bool а й структуру
		{[]int{4, 6}, 10},
		{[]int{11, 12}, 23},
	}

	for _, v := range tests {
		got := sum(v.xi[0], v.xi[1])
		want := v.want
		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	}
}
