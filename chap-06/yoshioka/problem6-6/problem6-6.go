package main

import (
	"fmt"
	"math"
)

func BinarySearch(a, b, c int) float64 {
	left := 0.0
	right := 1000.0

	for {
		mid := left + (right-left)/2

		r := calc(a, b, c, mid)

		if math.Abs(r-100.0) < math.Pow10(-11) {
			break
		}

		if r > 100 {
			right = mid
		} else {
			left = mid
		}
	}
	return right

}

func calc(a, b, c int, t float64) float64 {
	return float64(a)*t + float64(b)*math.Sin(float64(c)*t*math.Pi)
}

func main() {
	r := BinarySearch(1, 1, 1)
	fmt.Println("result", r)

}
