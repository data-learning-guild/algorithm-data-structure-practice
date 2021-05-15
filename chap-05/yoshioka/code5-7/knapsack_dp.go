package main

import "fmt"

func main(){
    var n int
	fmt.Print("Number of item length >>>")
	fmt.Scan(&n)

    var w int
	fmt.Print("Limit of weight >>>")
	fmt.Scan(&w)

	weight := setIntSlice(n)
	value := setIntSlice(n)

    db := KnapsackDP(n, w, weight, value)

    fmt.Println((db))

}

func KnapsackDP(n int, w int, weight []int, value []int) [][]int{
    // initialize db
    db := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        db[i] = make([]int, w+1)
    }

    for i := 0; i < n; i++ {
        for j := 0; j <= w; j++ {
            // choose the item number i
            if j - weight[i] >= 0 {
                relaxation(&db[i+1][j], db[i][j-weight[i]]+value[i])
            }
            // not choose
            relaxation(&db[i+1][j], db[i][j])
        }
    }
    return db
}



func relaxation(a *int, b int){
	if *a < b {
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
