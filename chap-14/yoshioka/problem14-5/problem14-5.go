package main

import (
	"fmt"
	"math"
)

const maxint int = math.MaxInt32

type Dequeue []int

func (queue *Dequeue) PushFront(i int) {
	*queue = append([]int{i}, (*queue)...)
}

func (queue *Dequeue) PushBack(i int) {
	*queue = append(*queue, i)
}

func (queue *Dequeue) Dequeue() int {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

func main() {
	k := 79992

	dist := make([]int, k)
	for i := 0; i < k; i++ {
		dist[i] = maxint
	}

	var que Dequeue

	dist[1] = 1
	que.PushFront(1)

	for len(que) > 0 {
		v := que.Dequeue()

		v2 := (v * 10) % k
		if dist[v2] > dist[v] {
			dist[v2] = dist[v]
			que.PushFront(v2)
		}

		v2 = (v + 1) % k
		if dist[v2] > dist[v]+1 {
			dist[v2] = dist[v] + 1
			que.PushBack(v2)

		}
	}

	fmt.Println(dist[0])

}
