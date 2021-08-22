package main

import (
	"fmt"

	"../queue"
)

type Graph [][]int

var color []int

func bfs(g *Graph, s int) bool {
	n := len(*g)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	var que queue.IntQueue

	color[s] = 0
	que.Enqueue(s)

	for len(que) > 0 {
		v := que.Dequeue()
		cur := color[v]

		for _, x := range (*g)[v] {
			if color[x] != -1 {
				if color[x] == cur {
					return false
				}
				continue
			}

			cur = 1 - color[v]
			color[x] = cur
			que.Enqueue(x)
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
		if !bfs(&g, v) {
			result = false
		}
	}

	fmt.Println("result:", result)
}
