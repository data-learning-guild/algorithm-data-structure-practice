package main

import (
	"fmt"
)

var memo map[int]int

func fibo(N int) int {
	// base case
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}

	// check the memo
	if memo[N] != -1 {
		return memo[N]
	}

	// update memo
	memo[N] = fibo(N-1) + fibo(N-2)
	return memo[N]
}

func main() {
	// initialise the memo
	memo = make(map[int]int)
	for i := 0; i < 50; i++ {
		memo[i] = -1
	}

	_ = fibo(49)

	for i := 2; i < 50; i++ {
		fmt.Printf("item %d: %d\n", i, memo[i])
	}
}
