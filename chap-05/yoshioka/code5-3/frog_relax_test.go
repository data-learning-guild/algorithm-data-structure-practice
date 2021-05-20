package main

import (
	"fmt"
	"testing"
)

func TestDP(t *testing.T){
    n:=7
    h := []int{2,9,4,5,1,6,10}
    memo := make([]int, n)
    memo = DP(n, h, memo)
    fmt.Println(memo)
}