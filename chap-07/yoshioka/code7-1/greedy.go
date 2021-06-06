package main

import (
	"fmt"
	"sort"
)

func Greedy(target_value int, cands map[int]int, order []int) int {
	result := 0

	for _, v := range order {
		// in case of no limit
		add := target_value / v

		amount, _ := cands[v]

		// check the maximum
		if add > amount {
			add = amount
		}

		target_value -= v * add
		result += add
	}

	return result
}

func main() {
	t := 551
	cands := map[int]int{
		500: 1,
		100: 10,
		50:  2,
		10:  1,
		5:   5,
		1:   4,
	}

	order := make([]int, len(cands))
	idx := 0
	for k := range cands {
		order[idx] = k
		idx++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(order)))

	r := Greedy(t, cands, order)
	fmt.Println("result: ", r)
}
