package main

import "fmt"

type Graph [][]int

var color []int

func dfs(g *Graph, v int, cur int) bool {
	color[v] = cur
	for _, next_v := range (*g)[v] {
		if color[next_v] != -1 {
			if color[next_v] == cur {
				return false
			}
			continue
		}
		if !dfs(g, next_v, 1-cur) {
			return false
		}
	}
	return true
}
func main() {
	n := 5

	g := make(Graph, n)
	g[0] = []int{1, 3}
	g[1] = []int{0, 2, 4}
	g[2] = []int{1}
	g[3] = []int{0, 4}
	g[4] = []int{1, 3}

	color = make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	result := true
	for v := 0; v < n; v++ {
		if color[v] != -1 {
			continue
		}
		if !dfs(&g, v, 0) {
			result = false
		}
	}

	fmt.Println("result:", result)
}
