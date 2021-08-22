package main

import (
	"fmt"

	"./../queue"
)

type Graph [][]int

var color []int

func bfs(g *Graph, s int) bool {
	n := len(*g)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	que := make(queue.IntQueue, 0)

	color[0] = 0
	cur := color[0]
	que.Enqueue(0)

	for len(que) > 0 {
		v := que.Dequeue()

		for _, x := range (*g)[v] {
			if color[x] != -1 {
				if color[x] == cur {
					return false
				}

				continue
			}

			color[x] = 1 - dist[v]
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
