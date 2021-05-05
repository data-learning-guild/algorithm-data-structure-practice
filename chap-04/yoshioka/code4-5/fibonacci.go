package main

import "fmt"

func Fibonacci(N int) int {
	//base case
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}

	// recursive
	return Fibonacci(N-1) + Fibonacci(N-2)

}

func main() {
	fmt.Print(6, "->", Fibonacci(6))
}
