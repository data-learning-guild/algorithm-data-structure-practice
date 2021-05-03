package main

import "testing"

func TestSevenFiveThree1(t *testing.T) {
	input_slice := []int{735, 375, 357, 573, 537, 753}

	for _, v := range input_slice {
		if !run(v, 1) {
			t.Error("Worng result for ", v)
		}
	}
}

func TestSevenFiveThree2(t *testing.T) {
	input_slice := []int{735375, 337755, 357537, 535773, 533577777, 7753353}

	for _, v := range input_slice {
		if !run(v, 2) {
			t.Error("Worng result for ", v)
		}
	}
}
func TestSevenFiveThree3(t *testing.T) {
	input_slice := []int{735373575, 373537755, 357357537, 573535773, 537353577777, 7773553353}

	for _, v := range input_slice {
		if !run(v, 3) {
			t.Error("Worng result for ", v)
		}
	}
}

func run(input int, expect_output int) bool {
	result := SevenFiveThree(input)
	return expect_output == result
}
