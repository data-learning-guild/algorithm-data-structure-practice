package main

import "fmt"

func Subset_sum(i int, w int, a []int) bool {
	//base case
	if i == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	// case: not choose a[i-1]
	if Subset_sum(i-1, w, a) {
		return true
	}

	// case: choose a[i-1]
	if Subset_sum(i-1, w-a[i-1], a) {
		return true
	}

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
