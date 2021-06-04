package main

import (
	"fmt"
	"sort"
)

func SnukeFestival(array_1, array_2, array_3 []int, length int) int {
	sort.Ints(array_1)
	sort.Ints(array_2)
	sort.Ints(array_3)

	result := 0

	for i := 0; i < length; i++ {
		fix_num := array_2[i]
		lower_num_val := LowerBound(array_1, fix_num)
		higher_num_val := length - UpperBound(array_3, fix_num)
		result += lower_num_val * higher_num_val
	}
	return result
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

func UpperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	l := 3
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	c := []int{3, 4, 5}
	// result: 1,2,3; 1,2,4; 1,2,5; 1,3,4; 1,3,5; 1,4,5; 2,3,4; 2,3,5; 2,4,5; 3,4,5 = 10
	r := SnukeFestival(a, b, c, l)
	fmt.Println("result", r)
}
