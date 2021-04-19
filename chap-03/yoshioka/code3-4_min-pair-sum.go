package main

import "fmt"

const inf int = 2000000000

func main() {
	var n, k int
	fmt.Print("list length >>>")
	fmt.Scan(&n)

	fmt.Print("border value >>>")
	fmt.Scan(&k)

	setA := setIntSlice(n)
	setB := setIntSlice(n)

	min_value := inf
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sumValue := setA[i] + setB[j]
			if sumValue < k {
				continue
			}
			if sumValue < min_value {
				min_value = sumValue
			}

		}

	}

	fmt.Println("min pair sum value: ", min_value)

}

func setIntSlice(n int) []int {
	intSlice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("slice item[", i, "] >>>")
		fmt.Scan(&intSlice[i])
	}
	return intSlice
}
