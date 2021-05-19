package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("list length >>>")
	fmt.Scan(&n)

	height_slice := setIntSlice(n)
	memo := make([]int, n)

	memo = DP(n, height_slice, memo)

	fmt.Println("Result:", memo[n-1])

}

func DP(n int, height_slice []int,memo []int) []int {
	for i := 0; i < n; i++ {
		memo[i] = math.MaxInt64
	}
	memo[0] = 0

	for i := 0; i < n; i++ {
		if i + 1 < n{
			relaxation(&memo[i+1], memo[i]+calc_diff(height_slice[i], height_slice[i+1]))
		}	
		if i + 2 < n {
			relaxation(&memo[i+2], memo[i]+calc_diff(height_slice[i], height_slice[i+2]))
		}

	}
	return memo
}

func calc_diff(start int, end int) int {
	float_result := math.Abs(float64(start) - float64(end))
	return int(float_result)
}

func relaxation(a *int, b int){
	if *a > b {
		*a = b
	}
}

func setIntSlice(n int) []int {
	intSlice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("slice item[", i, "] >>>")
		fmt.Scan(&intSlice[i])
	}
	return intSlice
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
