package main

import (
	"fmt"
)

type Graph [][]int

var seen []bool

func dfs(g *Graph, v int) {
	seen[v] = true

	for _, next_v := range (*g)[v] {
		if seen[next_v] {
			continue
		}
		dfs(g, next_v)
	}
}

func main() {
	n := 4
	g := make(Graph, n)
	g[0] = []int{1}
	g[1] = []int{0}
	g[2] = []int{3}
	g[3] = []int{2}

	seen = make([]bool, n)

	var cnt int

	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		dfs(&g, i)
		cnt++
	}

	fmt.Println("Count:", cnt)
}
