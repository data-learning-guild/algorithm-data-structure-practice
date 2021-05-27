package main

import "fmt"

func SubsetDP(number_list []int, target_num int) bool {
	n := len(number_list)
	dp_memo := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp_memo[i] = make([]bool, target_num+1)
	}

	dp_memo[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= target_num; j++ {
			// if !dp_memo[i][j] {
			// 	continue
			// }
			if dp_memo[i][j] {
				dp_memo[i+1][j] = true // 1. same as last i
			}
			if dp_memo[i+1][j] && j+number_list[i] <= target_num { // 2. the number of (base-number + next-number)  is less then the target number
				dp_memo[i+1][j+number_list[i]] = true
			}
		}
	}
	return dp_memo[n][target_num]

}

func main() {
	a := []int{5, 3, 4}
	result := SubsetDP(a, 29)
	fmt.Println(result)
}
