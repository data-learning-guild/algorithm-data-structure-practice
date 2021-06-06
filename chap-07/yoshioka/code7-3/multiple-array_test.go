package main

import "testing"

func TestMultipleArray1(t *testing.T) {
	//n := 3
	a := []int{3, 2, 9}
	b := []int{5, 7, 4}

	r := MultipleArray(a, b)

	if r != 7 {
		t.Error("Wrong result", r)
	}
}

func TestMultipleArray2(t *testing.T) {
	//n := 3
	a := []int{3, 4, 5, 2, 5, 5, 9}
	b := []int{1, 1, 9, 6, 3, 8, 7}

	r := MultipleArray(a, b)

	if r != 22 {
		t.Error("Wrong result", r)
	}
}
