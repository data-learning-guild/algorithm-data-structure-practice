package main

import "fmt"

func main(){
    var s, t string
    fmt.Print("1st word >>>")
	fmt.Scan(&s)
    fmt.Print("2nd word >>>")
	fmt.Scan(&t)

    dp := EditDistance(s, t)
    fmt.Println("result", dp[len(s)][len(t)])
    
}

func EditDistance(s string, t string) [][]int{
    dp := make([][]int, len(s)+1)
    for i := 0; i < len(s); i++ {
        dp[i] = make([]int, len(t)+1)
    }

    dp[0][0] = 0

    for i := 0; i <= len(s); i++ {
        for j := 0; j <= len(t); j++ {
            //Edit
            if i > 0 && j > 0 {
                if s[i-1] == t[j-1] {
                    relaxation(&dp[i][j], dp[i-1][j-1])
                }else{
                    relaxation(&dp[i][j], dp[i-1][j-1]+1)
                }
            }
            //Delete
            if i>0 {
                relaxation(&dp[i][j], dp[i-1][j]+1)
            }
            //Insert
            if j>0 {
                relaxation(&dp[i][j], dp[i][j-1]+1)
            }
    
        }
    }
    return dp

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
