package main

import "fmt"

var counter int

func Subset_sum(i int, w int, a []int) bool {
	counter = 0

	// initialise the memo
	memo := make(map[int]map[int]int)
	for j := 0; j < i+1; j++ {
		memo[j] = make(map[int]int)
		for k := 0; k < w+1; k++ {
			memo[j][k] = -1
		}
	}
	result := run_subset_sum(i, w, a, memo)
	fmt.Println("calc count: ", counter)
	return result

}

func run_subset_sum(i int, w int, a []int, memo map[int]map[int]int) bool {
	//base case
	if i == 0 {
		if w == 0 {
			memo[i][w] = 1
			return true
		} else {
			memo[i][w] = 0
			return false
		}
	}

	if memo[i][w] != -1 {
		switch memo[i][w] {
		case 1:
			return true
		case 0:
			return false
		default:
			panic("Unknown switch condition")
		}
	}

	counter++

	// case: not choose a[i-1]
	if run_subset_sum(i-1, w, a, memo) {
		memo[i][w] = 1
		return true
	}

	// case: choose a[i-1]
	if run_subset_sum(i-1, w-a[i-1], a, memo) {
		memo[i][w] = 1
		return true
	}

	memo[i][w] = 0
	return false

}

func main() {
	var n int
	fmt.Print("amount of numbers >>>")
	fmt.Scan(&n)

	var w int
	fmt.Print("target sum value >>>")
	fmt.Scan(&w)

	a := setIntSlice(n)

	if Subset_sum(n, w, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
