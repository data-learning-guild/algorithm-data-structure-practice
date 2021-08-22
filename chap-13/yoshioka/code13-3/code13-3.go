package main

import (
	"fmt"

	"../queue"
)

type Graph [][]int

func BFS(g *Graph, s int) []int {
	n := len(*g)

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	var que queue.IntQueue

	dist[s] = 0
	que.Enqueue(s)

	for len(que) > 0 {
		v := que.Dequeue()

		for _, x := range (*g)[v] {
			if dist[x] != -1 {
				continue
			}
			dist[x] = dist[v] + 1
			que.Enqueue(x)
		}
	}
	return dist
}

func main() {
	g := make(Graph, 4)
	g[0] = []int{1, 3}
	g[1] = []int{0, 2}
	g[2] = []int{1, 3}
	g[3] = []int{2, 0}

	s := 0
	dist := BFS(&g, s)

	for i := 0; i < len(g); i++ {
		fmt.Println(s, "->", i, "=", dist[i])
	}

}
