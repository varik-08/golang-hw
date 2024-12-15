package main

func BinarySearch(slice []int, target int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := low + (high-low)/2
		if slice[mid] == target {
			return mid
		}
		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
