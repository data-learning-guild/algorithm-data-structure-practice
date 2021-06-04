package lowerbound

// LowerBound ...
// https://gist.github.com/olee12/928b602d85afec4d8d991e2938e74d11
func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
