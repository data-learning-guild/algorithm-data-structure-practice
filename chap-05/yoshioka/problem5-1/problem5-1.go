// Problem 5.1 Vacation
// N days: summer vacation
// a[i] : happiness type A
// b[i] : happiness type B
// c[i] : happiness type C

package main

import "fmt"

func Vacation(period_days int, value_array [][]int) int {
	// initialize the memo
	dp := make([][]int, period_days+1)
	for i := 0; i < period_days+1; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < period_days; i++ { //day
		for j := 0; j < 3; j++ { //last item
			for k := 0; k < 3; k++ { //next item
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+value_array[i][k])
			}
		}

	}
	result := 0
	for i := 0; i < 3; i++ {
		chmax(&result, dp[period_days][i])
	}
	return result
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func main() {
	n := 3
	a := make([][]int, n)
	a[0] = []int{1, 2, 3}
	a[1] = []int{1, 2, 3}
	a[2] = []int{1, 2, 3}

	result := Vacation(n, a)
	fmt.Println(result)
}
