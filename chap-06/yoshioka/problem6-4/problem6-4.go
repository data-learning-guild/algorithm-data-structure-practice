package main

import (
	"fmt"
	"math"
)

func AggressiveCows(array []int, pick_num int) int {
	left := 0
	right := math.MaxInt64
	length := len(array)

	for right-left > 1 {
		mid := left + (right-left)/2 // value of distance

		cnt := 1  // 1st point was selected
		prev := 0 // last selected index

		// try case of "mid" distance in greedy method
		for i := 0; i < length; i++ {
			if array[i]-array[prev] >= mid {
				cnt++
				prev = i
			}
		}

		// judgement
		if cnt >= pick_num { // can be make in bigger distance
			left = mid
		} else {
			right = mid
		}
	}
	return left

}

func main() {
	a := []int{2, 4, 6, 8}
	m := 3
	r := AggressiveCows(a, m)
	fmt.Println("result", r)

}
