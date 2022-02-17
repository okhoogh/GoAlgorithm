package main

import "fmt"

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lengthTmp := length - i
		HeapSortMax(arr, lengthTmp)
		arr[0], arr[lengthTmp-1] = arr[lengthTmp-1], arr[0]
		fmt.Println("processing:", arr)
	}
	return arr
}

func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		top := i
		if left < length && arr[left] > arr[i] {
			i = left
		}
		if right < length && arr[right] > arr[i] {
			i = right
		}
		if top != i {
			arr[top], arr[i] = arr[i], arr[top]
		}
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(HeapSortMax(arr, len(arr)))
	fmt.Println(HeapSort(arr))
}
