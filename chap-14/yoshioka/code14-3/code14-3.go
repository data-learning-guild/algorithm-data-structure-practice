package main

import (
	"fmt"
	"math"
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

	used := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = maxint
	}
	dist[start] = 0

	for i := 0; i < n; i++ {
		min_dist := maxint
		min_v := -1

		for v := 0; v < n; v++ {
			if !used[v] && dist[v] < min_dist {
				min_dist = dist[v]
				min_v = v
			}
		}

		if min_v == -1 {
			break
		}

		for _, e := range g[min_v] {
			chmin(&dist[e.to], dist[min_v]+e.weight)
		}
		used[min_v] = true
	}

	for i := 0; i < n; i++ {
		fmt.Println(i, dist[i])
	}

}
