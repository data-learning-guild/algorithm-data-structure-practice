package main

import "fmt"

type Graph [][]int

var dp []int
var g Graph

func rec(v int) int {
	if dp[v] != -1 {
		return dp[v]
	}

	res := 0
	for _, nv := range g[v] {
		chmax(&res, rec(nv)+1)
	}
	dp[v] = res
	return dp[v]
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	} else {
		return false
	}
}

func main() {
	n := 4
	g = make(Graph, n)
	g[0] = []int{1, 2}
	g[1] = []int{3}
	g[2] = []int{1, 3}
	g[3] = []int{}

	dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	res := 0
	for v := 0; v < n; v++ {
		chmax(&res, rec(v))
	}

	fmt.Println(res)

}
