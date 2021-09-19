package FordFulkerson

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

func (g *Graph) Add_edge(from, to, cap int){
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

        flow := ff.fodfs(g, e.to, t, IntMin(f, e.cap))

        if flow==0 {
            continue
        }
    
        g.run_flow(e, flow)

        return flow
    }
    return 0
}

func (ff FordFulkerson)Solve(g *Graph, s, t int)int{
    var res int

    for {
        ff.seen = make([]bool, (*g).size())
        flow := ff.fodfs(g, s,t,Inf)

        if flow ==0 {
            return res
        }
        res += flow

    }
    return 0
}

func IntMin(a, b int) int{
    if a < b {
        return a
    } else{
        return b
    }
}

const Inf = math.MaxInt32

func main(){
    n := 6
    // m := 9

    g := make(Graph, n)

    input := [][]int{
        {0,1,5},
        {1,2,4},
        {1,3,37},
        {2,5,9},
        {2,4,56},
        {0,3,5},
        {3,2,3},
        {3,4,9},
        {4,5,2},
    }

    for _,e := range input{
        u, v, c := e[0], e[1], e[2]
        g.Add_edge(u,v,c)
    }

    ff := FordFulkerson{}
    s := 0
    t := n-1
    r := ff.Solve(&g, s,t)
    fmt.Println(r)
 
}