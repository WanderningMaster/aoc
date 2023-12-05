package utils

func BinarySearch[T int | string](nums []T, target T) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch[T int | string](nums []T, target T, ptrLeft int, ptrRight int) int {
	if ptrLeft > ptrRight {
		return -1
	}
	middle := (ptrLeft + ptrRight) / 2
	if nums[middle] == target {
		return middle
	}
	if nums[middle] > target {
		return binarySearch(nums, target, ptrLeft, middle-1)
	}

	return binarySearch(nums, target, middle+1, ptrRight)
}
