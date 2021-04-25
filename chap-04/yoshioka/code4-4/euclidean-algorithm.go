package main

import "fmt"

func GCD(m int, n int) int {
	// base case
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}

func main() {
	fmt.Print("15, 225 ->", GCD(15, 225))

}
