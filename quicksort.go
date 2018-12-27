package main

import (
	"fmt"
)

func QuickSort(slice []int64, left int64, right int64) {
	if left > right {
		return
	}
	base, i, j := slice[left], left, right
	for i != j {
		for slice[j] >= base && j > i {
			j--
		}
		for slice[i] <= base && i < j {
			i++
		}
		if i < j {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[left], slice[i] = slice[i], base
	QuickSort(slice, left, i-1)
	QuickSort(slice, i+1, right)
}

func main() {
	a := []int64{3, 6, 1, 7, 9, 5, 2, 4, 8}
	QuickSort(a, 0, 8)
	fmt.Println(a)
}
