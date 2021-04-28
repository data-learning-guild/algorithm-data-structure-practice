package main

import (
	"reflect"
	"testing"
)

func TestRemove_even_nums(t *testing.T) {
	test_slice1 := []int{4, 6, 8}
	test_slice2 := []int{4, 6, 1}
	test_slice3 := []int{8, 16, 24}

	cnt1, result1 := Remove_even_nums(0, test_slice1)

	if !reflect.DeepEqual(result1, []int{2, 3, 4}) || cnt1 != 1 {
		t.Error("Worng result in 1st slice")
	}

	cnt2, result2 := Remove_even_nums(0, test_slice2)

	if !reflect.DeepEqual(result2, test_slice2) || cnt2 != 0 {
		t.Error("Worng result in 2nd slice")
	}

	cnt3, result3 := Remove_even_nums(0, test_slice3)

	if !reflect.DeepEqual(result3, []int{1, 2, 3}) || cnt3 != 3 {
		t.Error("Worng result in 3rd slice")
	}

}
