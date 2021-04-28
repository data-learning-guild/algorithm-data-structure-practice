package main

import (
	"fmt"
	"math"
)

func main() {
	//Set the number of list length
	var n int
	fmt.Print("list length >>>")
	fmt.Scan(&n)

	if n <= 2 {
		panic("Length should be longer than 2")
	}

	number_list := setIntSlice(n)
	result := Find_second_min_number(n, number_list)

	fmt.Printf("second smallest number : %v\n", result)

}

func Find_second_min_number(length int, number_list []int) int {
	most_min_number := math.MaxInt64
	second_min_number := math.MaxInt64

	//Find most and second minimum number
	for i := 0; i < length; i++ {
		if number_list[i] < second_min_number {
			if number_list[i] < most_min_number {
				second_min_number = most_min_number
				most_min_number = number_list[i]
			} else {
				second_min_number = number_list[i]
			}
		}
	}
	return second_min_number

}

func setIntSlice(n int) []int {
	intSlice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("slice item[", i, "] >>>")
		fmt.Scan(&intSlice[i])
	}
	return intSlice
}
