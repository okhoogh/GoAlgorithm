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

func CountSort(arr []int) []int {
	length := len(arr)
	max, _ := SelectGetMax(arr)
	countArr := make([]int, max+1)
	tmpArr := make([]int, length)
	for _, v := range arr {
		countArr[v]++
	}
	for i := 1; i < max+1; i++ {
		countArr[i] += countArr[i-1]
	}
	for _, v := range arr {
		tmpArr[countArr[v]-1] = v
		countArr[v]--
	}
	return tmpArr
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1, 2, 5, 5, 3, 4, 3, 2, 1, 100, 94}
	fmt.Println(CountSort(arr))
}
