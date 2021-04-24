package main

import "fmt"

const inf = 2000000

func main() {
	var n int
	fmt.Print("list length >>>")
	fmt.Scan(&n)

	var v int
	fmt.Print("search value >>>")
	fmt.Scan(&v)

	a := setIntSlice(n)

	//Search
	min_value := inf
	for i := 0; i < n; i++ {
		if a[i] < min_value {
			min_value = a[i]
		}
	}
	fmt.Println("min value: ", min_value)
}

func setIntSlice(n int) []int {
	intSlice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("slice item[", i, "] >>>")
		fmt.Scan(&intSlice[i])
	}
	return intSlice
}
