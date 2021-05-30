package main

func BinarySearch(key int, array []int) int {
	left := 0
	right := len(array) - 1

	for right >= left {
		mid := left + (right-left)/2
		if array[mid] == key {
			return mid
		} else if array[mid] > key {
			right = mid - 1
		} else if array[mid] < key {
			left = mid + 1
		}
	}
	return -1
}
