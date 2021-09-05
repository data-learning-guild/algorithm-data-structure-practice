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

	g := make(Graph, n)
	g[0] = []Edge{
		{to: 1, weight: 3},
		{to: 3, weight: 100},
	}
	g[1] = []Edge{
		{to: 2, weight: 50},
		{to: 3, weight: 57},
		{to: 4, weight: -4},
	}
	g[2] = []Edge{
		{to: 3, weight: -10},
		{to: 4, weight: -10},
		{to: 5, weight: 100},
	}
	g[3] = []Edge{
		// {to: 3, weight: -5},
	}
	g[4] = []Edge{
		{to: 2, weight: 57},
		{to: 3, weight: 25},
		{to: 5, weight: 8},
	}
	g[5] = []Edge{}

	var exist_negative_cycle bool
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = maxint
	}

	start := 0
	dist[start] = 0

	for i := 0; i < n; i++ {
		var update bool
		for v := 0; v < n; v++ {
			if dist[v] == maxint {
				continue
			}

			for _, e := range g[v] {
				if chmin(&dist[e.to], dist[v]+e.weight) {
					update = true
				}
			}
		}
		if !update {
			break
		}
		if i == n-1 && update {
			exist_negative_cycle = true
		}
	}
	if exist_negative_cycle {
		fmt.Println("Negative cycle")
	} else {
		for i := 0; i < n; i++ {
			fmt.Println(i, dist[i])
		}
	}

}
