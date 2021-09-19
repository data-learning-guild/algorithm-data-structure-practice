package main

import (
	"fmt"
	"math"
)

type Edge struct{
    rev int
    from int
    to int
    cap int
}

type Graph [][]Edge

func (g *Graph) size() int{
    return len(*g)
}

func (g *Graph) redge(e *Edge) *Edge{
    return &((*g)[(*e).to][(*e).rev])
}

func (g *Graph) run_flow(e *Edge, f int){
    e.cap -= f
    g.redge(e).cap += f
}

func (g *Graph) add_edge(from, to, cap int){
    fromrev := len((*g)[from])
    torev := len((*g)[to])

    (*g)[from] = append((*g)[from], Edge{torev, from, to, cap})
    (*g)[to] = append((*g)[to], Edge{fromrev, to, from,0}) 
}

type FordFulkerson struct{
    seen []bool
}

// f: min flow in the path of s to v
func (ff FordFulkerson) fodfs(g *Graph, v, t, f int) int{
    if v == t {
        return f
    }

    //DFS
    ff.seen[v] = true
    for i := 0; i < len((*g)[v]); i++ {
        e := &(*g)[v][i]
        if ff.seen[e.to] {
            continue
        }
        if e.cap==0 {
            continue 
        }

        flow := ff.fodfs(g, e.to, t, min(f, e.cap))

        if flow==0 {
            continue
        }
    
        g.run_flow(e, flow)

        return flow
    }
    return 0
}

func (ff FordFulkerson)solve(g *Graph, s, t int)int{
    var res int

    for {
        ff.seen = make([]bool, (*g).size())
        flow := ff.fodfs(g, s,t,_inf)

        if flow ==0 {
            return res
        }
        res += flow

    }
    return 0
}

func min(a, b int) int{
    if a < b {
        return a
    } else{
        return b
    }
}

const _inf = math.MaxInt32

func main(){
    // Create Graph
    n := 6
    m := 9

    input := [][]int{
        {0,1,5,1},
        {1,2,4,1},
        {1,3,37,1},
        {2,5,9,1},
        {2,4,56,1},
        {0,3,5,5},
        {3,2,2,1},
        {3,4,9,2},
        {4,5,2,2},
    }

    // Distance matrix & Cost matrix
    dist := make([][]int, n)
    cap := make([][]int,n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, n)
        cap[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dist[i][j] = _inf
            cap[i][j] = _inf
        }
        dist[i][i] = 0
    }

    for i := 0; i < m; i++ {
        u, v, d, c := input[i][0],input[i][1],input[i][2],input[i][3]
        dist[u][v] = d
        cap[u][v] = c
    }

    // Floyd-Warshall
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = dist[i][j]
        }
    }
    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
            }
        }
    }


    //Create graph
    g := make(Graph, n)
    s := 0
    t := n-1

    for u := 0; u < n; u++ {
        for v := 0; v < n; v++ {
            if u==v {
                continue
            }
            if dp[s][u] + dist[u][v] + dp[v][t] == dp[s][t] { // 距離テーブルとDPコストが一致するなら最短路を構成する辺
                g.add_edge(u,v,cap[u][v])
            }
        }
    }

    ff := FordFulkerson{}
    r := ff.solve(&g, s,t)
    fmt.Println(r)
 
}