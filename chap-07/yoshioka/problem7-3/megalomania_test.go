package main

import "testing"

func TestMegalomania1(t *testing.T) {
	// n := 5
	a := []int{2, 1, 1, 4, 3}
	b := []int{4, 9, 8, 9, 12}

	r := Megalomania(a, b)
	if !r {
		t.Error("Wrong result", r)
	}
}
func TestMegalomania2(t *testing.T) {
	// n := 5
	a := []int{334, 334, 334}
	b := []int{1000, 1000, 1000}

	r := Megalomania(a, b)
	if r {
		t.Error("Wrong result", r)
	}
}
