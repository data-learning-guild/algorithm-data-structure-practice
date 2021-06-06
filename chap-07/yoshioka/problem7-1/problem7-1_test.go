package main

import "testing"

func TestGreeddy1(t *testing.T) {
	//n := 3
	a := []int{3, 2, 9}
	b := []int{5, 7, 4}

	r := Greedy(a, b)
	// (2,4) (3,5) = 2

	if r != 2 {
		t.Error("Wrong result", r)
	}
}
func TestGreeddy2(t *testing.T) {
	//n := 3
	a := []int{3, 2, 6}
	b := []int{5, 7, 4}

	r := Greedy(a, b)
	// (2,4) (3,5) (6,7)= 3

	if r != 3 {
		t.Error("Wrong result", r)
	}
}

func TestGreeddy3(t *testing.T) {
	//n := 3
	a := []int{5, 7, 4}
	b := []int{3, 2, 9}

	r := Greedy(a, b)
	// (4,9)=1

	if r != 1 {
		t.Error("Wrong result", r)
	}
}
