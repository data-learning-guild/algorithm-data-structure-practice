package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Start int
	End   int
}

type Pairs []Pair

// 以下インタフェースを満たす
// https://qiita.com/Jxck_/items/fb829b818aac5b5f54f7
func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pairs) Less(i, j int) bool {
	// Sort by the end times
	return p[i].End < p[j].End
}

func IntervalScheduling(pairs Pairs) int {
	sort.Sort(pairs)

	res := 0
	current_end_time := 0
	for i := 0; i < len(pairs); i++ {
		if pairs[i].Start < current_end_time {
			continue
		}
		res++
		current_end_time = pairs[i].End

	}
	return res
}

func main() {
	var p Pairs = []Pair{
		{9, 16},
		{10, 12},
		{11, 15},
		{13, 19},
		{15, 18},
		{19, 23},
	}
	r := IntervalScheduling(p)
	fmt.Println("result", r)

}
