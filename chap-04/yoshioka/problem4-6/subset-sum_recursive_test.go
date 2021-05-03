package main

import "testing"

func TestSubset_sum(t *testing.T) {
	n := 4
	w := 14
	a := []int{3, 2, 6, 5}

	if !Subset_sum(n, w, a) {
		t.Error("Wrong result in 1st data set")
	}

	n = 5
	w = 19
	a = []int{2, 4, 6, 8, 10}

	if Subset_sum(n, w, a) {
		t.Error("Wrong result in 2nd data set")
	}

	n = 5
	w = 21
	a = []int{1, 3, 7, 10, 13}

	if !Subset_sum(n, w, a) {
		t.Error("Wrong result in 3rd data set")
	}
}
func TestSubset_sum_big(t *testing.T) {
	n := 1000
	w := 944

	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i*2 + 1
	}
	if !Subset_sum(n, w, a) {
		t.Error("Wrong result in big data set")
	}

}
