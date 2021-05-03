package main

import "testing"

func TestTribonacci0(t *testing.T) {
	input := 0
	result_bool := run(input, 0)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}

func TestTribonacci1(t *testing.T) {
	input := 1
	result_bool := run(input, 0)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}

func TestTribonacci2(t *testing.T) {
	input := 2
	result_bool := run(input, 1)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}

func TestTribonacci3(t *testing.T) {
	input := 3
	result_bool := run(input, 1)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}

func TestTribonacci4(t *testing.T) {
	input := 4
	result_bool := run(input, 2)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}

func TestTribonacci5(t *testing.T) {
	input := 5
	result_bool := run(input, 4)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}
func TestTribonacci10(t *testing.T) {
	input := 9
	result_bool := run(input, 44)
	if !result_bool {
		t.Error("Worng output for ", input)
	}
}

func run(input int, expect_output int) bool {
	result := Tribonacci(input)
	if result == expect_output {
		return true
	} else {
		return false
	}
}
