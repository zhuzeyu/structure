package main

import (
	"fmt"
)

func BinarySearch(slice []int64, key int64) int64 {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == key {
			return int64(mid)
		} else if slice[mid] > key {
			right = mid - 1
		} else if slice[mid] < key {
			left = mid + 1
		}
	}
	return -1
}

func SquareRoot(num int64) int64 {
	left, right := int64(1), num
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return mid
		} else if mid*mid > num {
			right = mid - 1
		} else if mid*mid < num {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	a := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(a, 5))
	fmt.Println(BinarySearch(a, 99))
	fmt.Println(SquareRoot(25))
	fmt.Println(SquareRoot(26))
}
