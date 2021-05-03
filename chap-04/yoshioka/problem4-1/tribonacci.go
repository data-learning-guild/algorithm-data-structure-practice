package main

import "fmt"

func Tribonacci(N int) int {
	if (N == 0) || (N == 1) {
		return 0
	} else if N == 2 {
		return 1
	} else {
		return Tribonacci(N-1) + Tribonacci(N-2) + Tribonacci(N-3)
	}
}

func main() {
	fmt.Println(Tribonacci(7))
}
