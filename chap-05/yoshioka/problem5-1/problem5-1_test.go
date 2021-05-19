package main

import (
	"testing"
)

func testVacation(t *testing.T) {
	n := 3
	a := make([][]int, n)
	a[0] = []int{1, 1, 1}
	a[1] = []int{2, 2, 2}
	a[2] = []int{3, 3, 3}

	result := Vacation(n, a)
	if result != 8 {
		t.Error("Worng result")
	}
}
