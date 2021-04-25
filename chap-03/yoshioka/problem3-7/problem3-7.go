package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Println("Enter the number >>>")
	fmt.Scan(&n)

	result := Sum_string_numbers(n)
	fmt.Println("Result:", result)

}

func Sum_string_numbers(str_number string) int {
	length := len(str_number)
	sum := 0

	for bit := 0; bit < (1 << (length - 1)); bit++ {
		last_idx := 0
		for i := 0; i < (length - 1); i++ {

			if (bit & (1 << i)) > 0 {
				tmp, _ := strconv.Atoi(str_number[last_idx : i+1]) // sum until the last index to current index
				sum += tmp
				last_idx = i + 1

			}
		}
		tmp, _ := strconv.Atoi(str_number[last_idx:]) // add the remaining portions
		sum += tmp

	}
	return sum

}
