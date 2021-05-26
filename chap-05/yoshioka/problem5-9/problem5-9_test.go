package main

import "testing"

func TestBinarySearchTree1(t *testing.T) {
	a := []int{1, 2, 3}
	r := BinarySearchTree(a)
	correct := 9
	if r != correct {
		t.Error("Worng output in 1st test", r, correct)
	}
}

func TestBinarySearchTree2(t *testing.T) {
	a := []int{1, 2}
	r := BinarySearchTree(a)
	correct := 3
	if r != correct {
		t.Error("Worng output in 2nd test", r, correct)
	}
}

func TestBinarySearchTree3(t *testing.T) {
	a := []int{2, 5, 4}
	r := BinarySearchTree(a)
	correct := 18
	if r != correct {
		t.Error("Worng output in 3rd test", r, correct)
	}
}
func TestBinarySearchTree4(t *testing.T) {
	a := []int{1, 2, 3, 4}
	r := BinarySearchTree(a)
	correct := 19
	if r != correct {
		t.Error("Worng output in 4th test", r, correct)
	}
}
func TestBinarySearchTree5(t *testing.T) {
	a := []int{4, 3, 2}
	r := BinarySearchTree(a)
	correct := 14
	if r != correct {
		t.Error("Worng output in 5th test", r, correct)
	}
}
