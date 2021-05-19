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

	dp_memo := make([]int, n)
	for i := 0; i < n; i++ {
		dp_memo[i] = math.MaxInt64
	}
	dp_memo[0] = 0

	for i := 1; i < n; i++ {
		if i == 1 {
			dp_memo[i] = calc_diff(height_slice[i], height_slice[i-1])
		} else {
			dp_memo[i] = Min(dp_memo[i-1]+calc_diff(height_slice[i], height_slice[i-1]), dp_memo[i-2]+calc_diff(height_slice[i], height_slice[i-2]))
		}
	}

	fmt.Println("Result:", dp_memo[n-1])

}

func calc_diff(start int, end int) int {
	float_result := math.Abs(float64(start) - float64(end))
	return int(float_result)
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
