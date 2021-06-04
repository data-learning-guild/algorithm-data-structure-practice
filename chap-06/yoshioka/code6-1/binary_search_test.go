package main

import "testing"

const N int = 8

var array = []int{3, 5, 8, 10, 14, 17, 21, 39}

func TestBinerySearch1(t *testing.T) {
	r := BinarySearch(10, array)
	if r != 3 {
		t.Error()
	}
}
func TestBinerySearch2(t *testing.T) {
	r := BinarySearch(3, array)
	if r != 0 {
		t.Error()
	}
}
func TestBinerySearch3(t *testing.T) {
	r := BinarySearch(39, array)
	if r != 7 {
		t.Error()
	}
}

func TestBinerySearch4(t *testing.T) {
	r := BinarySearch(100, array)
	if r != -1 {
		t.Error()
	}
}

func TestBinerySearch5(t *testing.T) {
	r := BinarySearch(9, array)
	if r != -1 {
		t.Error()
	}
}
