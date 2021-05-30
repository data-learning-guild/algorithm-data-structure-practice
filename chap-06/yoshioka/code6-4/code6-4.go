package main

import (
	"fmt"
	"math"
	"sort"
)

func min_sum(slice_1, slice_2 []int, length, target_num int) int {
	sort.Ints(slice_1)
	sort.Ints(slice_2)
	min_value := math.MaxInt64
	for i := 0; i < length; i++ { //O(N)
		lb := LowerBound(slice_2, target_num-slice_1[i]) //O(logN)
		lb_num := slice_2[lb]
		if slice_1[i]+lb_num < min_value {
			min_value = slice_1[i] + lb_num
		}
	}
	return min_value
}

func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	n := 3
	k := 5
	a := []int{3, 2, 1}
	b := []int{4, 3, 2}

	r := min_sum(a, b, n, k)
	fmt.Println("result", r)

}
