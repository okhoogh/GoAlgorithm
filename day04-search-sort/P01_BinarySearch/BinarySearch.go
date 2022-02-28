package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	low := make([]int, 0)
	mid := make([]int, 0)
	high := make([]int, 0)
	target := arr[rand.Intn(length)]
	for i := 0; i < length; i++ {
		if arr[i] < target {
			low = append(low, arr[i])
		} else if arr[i] > target {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	left, right := QuickSort(low), QuickSort(high)
	return append(append(left, mid...), right...)
}

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if val < arr[mid] {
			right = mid - 1
		} else if val > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 19, 4, 8, 3, 5, 4, 6, 19, 0}
	fmt.Println("arr:", arr)
	fmt.Println("sorted:", QuickSort(arr))
	arr = QuickSort(arr)
	index := BinarySearch(arr, 19)
	if index != -1 {
		fmt.Println("search:", index)
	} else {
		fmt.Println("no data")
	}
}
