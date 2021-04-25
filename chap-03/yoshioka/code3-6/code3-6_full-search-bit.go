package main

import "fmt"

func main() {
	var n, w int
	fmt.Println("list length >>>")
	fmt.Scan(&n)

	fmt.Println("target number >>>")
	fmt.Scan(&w)

	candNumbers := setIntSlice(n)

	var exist bool
	for bit := 0; bit < (1 << n); bit++ {
		sum := 0
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) > 0 {
				sum += candNumbers[i]
				fmt.Println(bit, ", ", i, ", ", candNumbers[i], ", ", sum)
			}
		}
		if sum == w {
			exist = true
		}
	}
	if exist {
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
