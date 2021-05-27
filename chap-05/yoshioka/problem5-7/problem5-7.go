package main

import "fmt"

func LCS(string1, string2 string) int {
	len1 := len(string1)
	len2 := len(string2)

	//initialize
	dp_memo := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp_memo[i] = make([]int, len2+1)
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if string1[i] == string2[j] {
				dp_memo[i+1][j+1] = max(dp_memo[i+1][j+1], dp_memo[i][j]+1)
			}

			//Vertical and horizontal transitions should be considered in all cases
			dp_memo[i+1][j+1] = max(dp_memo[i+1][j+1], dp_memo[i+1][j])
			dp_memo[i+1][j+1] = max(dp_memo[i+1][j+1], dp_memo[i][j+1])
		}
	}

	return dp_memo[len1][len2]

}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	s1 := "abcde"
	s2 := "acbef"

	r := LCS(s1, s2)
	fmt.Println("result:", r)
}
