package main

import (
	"testing"
)

func TestSum_string_numbers1(t *testing.T) {
	str_num1 := "125"
	result1 := Sum_string_numbers(str_num1)
	if result1 != 176 { // 125 + (1 + 25) + (12 + 5) + (1 + 2 +5)
		t.Error("Worng result for 125")
	}

}

func TestSum_string_numbers2(t *testing.T) {
	str_num1 := "3333"
	result1 := Sum_string_numbers(str_num1)
	if result1 != 4200 {
		t.Error("Worng result for 3333")
	}

}
