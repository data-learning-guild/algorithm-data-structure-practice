package main

import "fmt"

type Graph [][]int

var seen []bool
var order []int

func rec(g *Graph, v int) {
	seen[v] = true
	for _, next_v := range (*g)[v] {
		if seen[next_v] {
			continue
		}
		rec(g, next_v)
	}
	order = append(order, v)
}

func main() {
	n := 8

	g := make(Graph, n)
	g[0] = []int{5}
	g[1] = []int{6, 3}
	g[2] = []int{5, 7}
	g[3] = []int{0, 7}
	g[4] = []int{1, 6, 2}
	g[5] = []int{}
	g[6] = []int{7}
	g[7] = []int{0}

	seen = make([]bool, n)
	order = make([]int, 0)

	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		rec(&g, i)
	}

	for o := range order {
		fmt.Println("->", order[len(order)-1-o])
	}
}
