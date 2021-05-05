package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	result1 := Fibonacci(6)
	if result1 != 8 {
		t.Error("Worng result in input: ", 6)
	}

	result2 := Fibonacci(20)
	if result2 != 6765 {
		t.Error("Worng result in input: ", 20)
	}

}
