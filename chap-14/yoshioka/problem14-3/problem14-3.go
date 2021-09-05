package main

import "fmt"

type Graph [][]int
type Pair struct {
	id     int
	parity int
}

type PairQueue []Pair

func (queue *PairQueue) Enqueue(p Pair) {
	*queue = append(*queue, p)
}

func (queue *PairQueue) Dequeue() Pair {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

var n int
var g Graph

func solve(s, t int) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			dist[i][j] = -1
		}
	}

	dist[s][0] = 0
	var que PairQueue
	que.Enqueue(Pair{s, 0})

	for len(que) > 0 {
		p := que.Dequeue()
		v := p.id
		parity := p.parity

		for _, nv := range g[v] {
			np := (parity + 1) % 3
			if dist[nv][np] == -1 {
				dist[nv][np] = dist[v][parity] + 1
				que.Enqueue(Pair{nv, np})
			}

		}

	}
	if dist[t][0] == -1 {
		return -1
	} else {
		return dist[t][0] / 3
	}

}

func main() {
	n = 4
	g = make(Graph, n)
	g[0] = []int{1}
	g[1] = []int{2}
	g[2] = []int{3}
	g[3] = []int{0}

	fmt.Println(solve(0, 2))
}
