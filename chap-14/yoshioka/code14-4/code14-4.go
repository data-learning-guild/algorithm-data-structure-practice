package main

import (
	"fmt"
	"math"

	"../heap"
)

const maxint int = math.MaxInt32

type Edge struct {
	to     int
	weight int
}

type Graph [][]Edge

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	} else {
		return false
	}
}

func main() {
	n := 6
	start := 0

	g := make(Graph, n)
	g[0] = []Edge{
		{to: 1, weight: 3},
		{to: 2, weight: 5},
	}
	g[1] = []Edge{
		{to: 3, weight: 12},
		{to: 2, weight: 4},
	}
	g[2] = []Edge{
		{to: 3, weight: 9},
		{to: 4, weight: 4},
	}
	g[3] = []Edge{
		{to: 5, weight: 2},
	}
	g[4] = []Edge{
		{to: 3, weight: 7},
		{to: 5, weight: 8},
	}
	g[5] = []Edge{}

	// Dijkstra dist
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = maxint
	}
	dist[start] = 0

	//Heap
	que := make(heap.Heap, n)
	que.Push(heap.Pair{dist[start], start})

	for len(que) > 0 {
		t := que.Top()
		v := t.Idx
		d := t.Dist

		que.Pop()

		if d > dist[v] {
			continue
		}

		for _, e := range g[v] {
			if chmin(&dist[e.to], dist[v]+e.weight) {
				que.Push(heap.Pair{dist[e.to], e.to})
			}
		}
	}

	for v := 0; v < n; v++ {
		if dist[v] < maxint {
			fmt.Println(v, dist[v])
		} else {
			fmt.Println(v, "INF")
		}
	}

}
