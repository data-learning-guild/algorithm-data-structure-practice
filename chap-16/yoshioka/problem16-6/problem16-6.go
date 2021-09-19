package main

import (
	"fmt"

	"../FordFulkerson"
)


func main(){
    //input
    n :=2
    a := []int{-1000,100000}
    size := len(a) + 2 //dummy node

    //graph
    g := make(FordFulkerson.Graph, size)
    var offset int
    for i,v := range a{
        idx := i+1
        if v >= 0{
            g.Add_edge(0,idx,0)
            g.Add_edge(idx, size-1, v)
            offset += v
        }else{
            g.Add_edge(0,idx,-v)
            g.Add_edge(idx, size-1,0)
        }
    }

    for i := 1; i < n+1; i++ {
        for j := 1; j < n+1; j++ {
            if j % i == 0 {
                g.Add_edge(i, j, FordFulkerson.Inf)
            }
        }
    }

    ff := FordFulkerson.FordFulkerson{}
    res := ff.Solve(&g,0, size-1)
    fmt.Println(offset - res)

}