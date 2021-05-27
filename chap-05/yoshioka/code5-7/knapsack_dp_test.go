package main

import (
	"testing"
)

func TestDP(t *testing.T){
    n:=6
    w:= 15
    weight := []int{2,1,3,2,1,5}
    value := []int{3,2,6,1,3,85}
    result := KnapsackDP(n,w,weight,value)

    if result[n][w] != 100{
        t.Error("Worng result")
    }

}