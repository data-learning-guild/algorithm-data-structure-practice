package main

import (
	"testing"
)

func TestFind_second_min_number(t *testing.T) {
	test_list1 := []int{1, 2, 3}
	result1 := run_func(len(test_list1), test_list1, 2)
	if !result1 {
		t.Error("Worng result in 1st test")
	}

	test_list2 := []int{11, 442, 24, 77, 9, 2}
	result2 := run_func(len(test_list2), test_list2, 9)
	if !result2 {
		t.Error("Worng result in 1st test")
	}

}

func run_func(input1 int, input2 []int, output int) bool {
	result := Find_second_min_number(input1, input2)
	if result != output {
		return false
	} else {
		return true

	}
}
