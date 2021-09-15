package main

import (
	"sort"

	"../unionfind"
)

type node_pair struct {
	start int
	end   int
}
type Edge struct {
	weight int
	nodes  node_pair
}

type Edges []Edge

func (e Edges) Len() int {
	return len(e)
}
func (e Edges) Less(i, j int) bool {
	return e[i].weight < e[j].weight
}
func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	//input
	n := 8
	m := 10
	edges := make(Edges, m)
	edges[0] = Edge{weight: 9, nodes: node_pair{4, 2}}
	edges[1] = Edge{weight: 2, nodes: node_pair{4, 6}}
	edges[2] = Edge{weight: 4, nodes: node_pair{4, 1}}
	edges[3] = Edge{weight: 7, nodes: node_pair{6, 7}}
	edges[4] = Edge{weight: 8, nodes: node_pair{1, 3}}
	edges[5] = Edge{weight: 5, nodes: node_pair{2, 7}}
	edges[6] = Edge{weight: 10, nodes: node_pair{2, 5}}
	edges[7] = Edge{weight: 3, nodes: node_pair{7, 0}}
	edges[8] = Edge{weight: 5, nodes: node_pair{3, 0}}
	edges[9] = Edge{weight: 6, nodes: node_pair{0, 5}}

	//sort by weight
	sort.Sort(edges)

	//Kruskal's algorithm
	res := 0
	var uf unionfind.UnionFind
	uf.Ini(n)

	for i := 0; i < m; i++ {
		w := edges[i].weight
		u, v := edges[i].nodes.start, edges[i].nodes.end

		if uf.Issame(u, v) {
			continue
		}

		res += w
		uf.Unite(u, v)

	}

}
