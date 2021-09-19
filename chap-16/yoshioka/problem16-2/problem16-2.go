package main

import (
	"fmt"

	"../FordFulkerson"
)


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
            dist[i][j] = FordFulkerson.Inf
            cap[i][j] = FordFulkerson.Inf
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
                dp[i][j] = FordFulkerson.IntMin(dp[i][j], dp[i][k]+dp[k][j])
            }
        }
    }


    //Create graph
    g := make(FordFulkerson.Graph, n)
    s := 0
    t := n-1

    for u := 0; u < n; u++ {
        for v := 0; v < n; v++ {
            if u==v {
                continue
            }
            if dp[s][u] + dist[u][v] + dp[v][t] == dp[s][t] { // 距離テーブルとDPコストが一致するなら最短路を構成する辺
                g.Add_edge(u,v,cap[u][v])
            }
        }
    }

    ff := FordFulkerson.FordFulkerson{}
    r := ff.Solve(&g, s,t)
    fmt.Println(r)
 
}