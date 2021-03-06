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
			if !dp_memo[i][j] {
				continue
			}
			dp_memo[i+1][j] = true
			if j+number_list[i] <= target_num {
				dp_memo[i+1][j+number_list[i]] = true
			}
		}
	}
	return dp_memo[n][target_num]

}

func main() {
	a := []int{1, 3, 4}
	result := SubsetDP(a, 7)
	fmt.Println(result)
}
