package main

import "fmt"

func SubsetDP(number_list []int, target_num int) int {
	n := len(number_list)
	// initialize memo
	dp_memo := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp_memo[i] = make([]bool, target_num+1)
	}

	dp_memo[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= target_num; j++ { //find until target num
			if !dp_memo[i][j] { //skip the case of did not select
				continue
			}
			dp_memo[i+1][j] = true //select the current num
			if j+number_list[i] <= target_num {
				dp_memo[i+1][j+number_list[i]] = true // select next num
			}
		}
	}

	result := 0
	for i := 1; i <= target_num; i++ {
		if dp_memo[n][i] {
			result++ // count the numbers that can be issued
		}
	}
	return result

}

func main() {
	a := []int{1, 2, 3, 4, 5}
	result := SubsetDP(a, 15)
	fmt.Println(result)
}
