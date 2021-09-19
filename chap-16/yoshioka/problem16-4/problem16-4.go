package main

import (
	"fmt"

	"../FordFulkerson"
	"../GCD"
)


func main(){
    //input
    a := []int{2,3,5,7,9,11,13,15,17,29}
    b := []int{4,6,10,14,18,22,26,30,34,38}
    n := len(a) + len(b) + 2 // dummy node

    g:= make(FordFulkerson.Graph, n)
    
    a_idx := make(map[int]int, len(a))
    for i,v := range a{
        idx := i + 1 // push one for dummy node
        g.Add_edge(0,idx,1)
        a_idx[idx] = v
    }
    b_idx := make(map[int]int, len(b))
    for i,v := range b{
        idx := i + len(a) + 1
        g.Add_edge(idx, n-1,1)
        b_idx[idx] = v
    }

    for ak, av := range a_idx{
        for bk, bv := range b_idx{
            if GCD.GCD(av, bv) > 1{
                g.Add_edge(ak, bk, 1)
            }
        }
    }

    ff := FordFulkerson.FordFulkerson{}
    fmt.Println(ff.Solve(&g, 0, n-1))


}