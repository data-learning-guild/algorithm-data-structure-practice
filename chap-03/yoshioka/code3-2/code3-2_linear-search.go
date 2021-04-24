package main

import "fmt"

func main() {
	var n int
	fmt.Print("list length >>>")
	fmt.Scan(&n)

	var v int
	fmt.Print("search value >>>")
	fmt.Scan(&v)

	a := setIntSlice(n)

	//Search
	found_id := -1
	for i := 0; i < n; i++ {
		if a[i] == v {
			found_id = i
			break
		}
	}
	if found_id > 0 {
		fmt.Println("Yes", found_id)
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
