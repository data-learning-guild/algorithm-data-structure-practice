package main

import "sort"

func Greedy(array1, array2 []int) int {
	n := len(array1)
	if n != len(array2) {
		panic("Wrong input argument")
	}

	sort.Ints(array1) //O(NlogN)
	sort.Ints(array2)

	base_idx := 0
	for i := 0; i < n; i++ {
		a := array1[base_idx]
		b := array2[i]
		if a < b {
			base_idx++
		}
	}
	return base_idx
}
