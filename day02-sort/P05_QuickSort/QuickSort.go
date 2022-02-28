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
	splitNum := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	for i := 0; i < length; i++ {
		if splitNum > arr[i] {
			low = append(low, arr[i])
		} else if splitNum < arr[i] {
			high = append(high, arr[i])
		} else {
			mid = append(mid, splitNum)
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}

func QuickSortRandom(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	splitNum := arr[rand.Intn(length)]
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	for i := 0; i < length; i++ {
		if splitNum > arr[i] {
			low = append(low, arr[i])
		} else if splitNum < arr[i] {
			high = append(high, arr[i])
		} else {
			mid = append(mid, splitNum)
		}
	}
	low, high = QuickSortRandom(low), QuickSortRandom(high)
	return append(append(low, mid...), high...)
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(QuickSort(arr))
	fmt.Println(QuickSortRandom(arr))
}
