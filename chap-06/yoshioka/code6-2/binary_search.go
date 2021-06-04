package main

func p(x int) bool {
	// write the condition
	return false
}

func BinarySearch() int {
	var left, right int //p(left)=false, p(right)=true

	for right-left > 1 {
		mid := left + (right-left)/2
		if p(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
