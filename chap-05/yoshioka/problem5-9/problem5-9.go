package main

import (
	"fmt"
)

func BinarySearchTree(size_list []int) int {
	length := len(size_list)

	//Prepare for Cumulative sum
	cum_sum := make([]int, length+1)
	for i := 0; i < length; i++ {
		cum_sum[i+1] = cum_sum[i] + size_list[i]
	}

	//initialize
	dp_memo := make([][]int, length+1)
	for i := 0; i < len(dp_memo); i++ {
		dp_memo[i] = make([]int, length+1)
		for j := 0; j < len(dp_memo); j++ {
			dp_memo[i][j] = 100000000
		}
	}

	// initial status
	for i := 0; i < length; i++ {
		dp_memo[i][i+1] = 0
	}

	//start updating
	// not works
	for j := 2; j <= length; j++ {
		for i := 0; i+1 < j; i++ {

			for k := i + 1; k < j; k++ {
				chmin(&dp_memo[i][j], dp_memo[i][k]+dp_memo[k][j]+(cum_sum[j]-cum_sum[i]))
			}

		}
	}

	// answer
	// for bet := 2; bet <= length; bet++ {
	// 	for i := 0; i+bet <= length; i++ {
	// 		j := i + bet
	// 		for k := i + 1; k < j; k++ {
	// 			chmin(&dp_memo[i][j], dp_memo[i][k]+dp_memo[k][j]+(cum_sum[j]-cum_sum[i]))
	// 		}
	// 	}
	// }

	return dp_memo[0][length]
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func main() {
	a := []int{1, 2, 3}
	r := BinarySearchTree(a)
	fmt.Println("result", r)
}
