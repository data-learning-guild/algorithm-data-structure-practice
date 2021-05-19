package main

import "testing"

func TestSubsetDP(t *testing.T) {
	a := []int{1, 3, 5}
	r := SubsetDP(a, 4)
	if r != 3 {
		t.Error("worng result in ", 1)
	}
}
func TestSubsetDP2(t *testing.T) {
	a := []int{1, 2, 3}
	r := SubsetDP(a, 5)
	if r != 5 {
		t.Error("worng result in 2")
	}
}
func TestSubsetDP3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	r := SubsetDP(a, 3)
	if r != 3 {
		t.Error("worng result in 3")
	}
}
