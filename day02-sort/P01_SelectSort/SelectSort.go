package main

import (
	"errors"
	"fmt"
)

func SelectGetMax(arr []int) (int, error) {
	size := len(arr)
	if arr == nil || size == 0 {
		return -1, errors.New("array is empty")
	}
	if size == 1 {
		return arr[0], nil
	}
	Max := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] > Max {
			Max = arr[i]
		}
	}
	return Max, nil
}

func SelectSort(arr []int) ([]int, error) {
	size := len(arr)
	if size == 1 {
		return arr, nil
	}
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr, nil
}

func main111() {
	arr := []int{5, 4, 3, 2, 1}
	max, _ := SelectGetMax(arr)
	fmt.Println(max)

	sort, _ := SelectSort(arr)
	fmt.Println(sort)
}
