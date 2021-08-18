package main

import "./../queue"

type Graph [][]int

func search(g *Graph, s int) {
	n := len(*g)

	seen := make([]bool, n)
	todo := make(queue.IntQueue, 0)

	//ini
	seen[s] = true
	todo.Enqueue(s)

	//while todo is empty
	for len(todo) > 0 {
		v := todo.Dequeue()

		for _, x := range (*g)[v] {
			if seen[x] {
				continue
			}

			seen[x] = true
			todo.Enqueue(x)
		}
	}

}

func main() {
	g := make(Graph, 3)
	g[0] = []int{1}
	g[1] = []int{0, 2}
	g[2] = []int{1}

	search(&g, 0)

}
