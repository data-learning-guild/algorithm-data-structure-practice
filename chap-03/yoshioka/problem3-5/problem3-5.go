package main

import (
	"fmt"
)

func main() {
	//Set the number of list length
	var n int
	fmt.Print("list length >>>")
	fmt.Scan(&n)

	number_list := setIntSlice(n)
	cnt, result := Remove_even_nums(0, number_list)

	fmt.Println(cnt)
	fmt.Println(result)

}

func Remove_even_nums(cnt int, input_slice []int) (int, []int) {
	result_slice := make([]int, len(input_slice))

	for i := 0; i < len(input_slice); i++ {
		n := input_slice[i]
		if n%2 == 0 {
			result_slice[i] = n / 2
		} else {
			return cnt, input_slice
		}

	}
	cnt += 1
	return Remove_even_nums(cnt, result_slice)

}

func setIntSlice(n int) []int {
	intSlice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("slice item[", i, "] >>>")
		fmt.Scan(&intSlice[i])
	}
	return intSlice
}
