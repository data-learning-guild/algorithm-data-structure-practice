package main

import (
	"fmt"

	"../queue"
)

type DiGraph [][]int

var order []int
var deg []int

func BFS(g *DiGraph) {
	//ini
	var que queue.IntQueue
	for i := 0; i < len(*g); i++ {
		if deg[i] == 0 {
			que.Enqueue(i)
		}
	}

	//while que is empty
	for len(que) > 0 {
		v := que.Dequeue()
		order = append(order, v)

		for _, x := range (*g)[v] { //TODO change to rev
			deg[x]--

			if deg[x] == 0 {
				que.Enqueue(x)
			}
		}
	}

}

func main() {
	n := 8

	g := make(DiGraph, n)
	g[0] = []int{5}
	g[1] = []int{6, 3}
	g[2] = []int{5, 7}
	g[3] = []int{0, 7}
	g[4] = []int{1, 6, 2}
	g[5] = []int{}
	g[6] = []int{7}
	g[7] = []int{0}

	// seen = make([]bool, n)
	order = make([]int, 0)

	deg = make([]int, n)
	for i := 0; i < n; i++ {
		for _, e := range g[i] {
			deg[e]++
		}
	}

	BFS(&g)

	for _, o := range order {
		fmt.Println("->", o)
	}
}
