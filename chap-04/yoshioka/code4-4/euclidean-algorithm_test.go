package main

import "testing"

func TestGCD(t *testing.T) {
	result1 := GCD(51, 15)
	if result1 != 3 {
		t.Error("Error in 51, 15")
	}
	result2 := GCD(15, 51)
	if result2 != 3 {
		t.Error("Error in 15, 51")
	}

}
