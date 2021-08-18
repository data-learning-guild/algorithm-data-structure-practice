package main

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
	n := 3
	g := make(Graph, n)
	g[0] = []int{1}
	g[1] = []int{0, 2}
	g[2] = []int{1}

	seen = make([]bool, n)

	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		dfs(&g, v)
	}

}
