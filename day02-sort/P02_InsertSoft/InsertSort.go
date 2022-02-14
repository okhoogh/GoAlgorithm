package main

import "fmt"

func InsertTest(arr []int) []int {
	j := 3
	backup := arr[j]
	for ; j > 0 && backup < arr[j-1]; j-- {
		arr[j] = arr[j-1]
	}
	arr[j] = backup
	return arr
}

func InsertSort1(arr []int) []int {
	size := len(arr)
	if size == 1 {
		return arr
	}
	for i := 1; i < size; i++ {
		backup := arr[i]
		j := i
		for ; j > 0 && backup < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = backup
	}
	return arr
}

func InsertSort2(arr []int) []int {
	size := len(arr)
	if size == 1 {
		return arr
	}
	for i := size - 2; i >= 0; i-- {
		j := i
		for ; j < size-1 && backup > arr[j+1]; j++ {
			arr[j+1] = arr[j]
		}
		arr[i] = backup
	}
	return arr
}

func main() {
	arr := []int{100, 19, 29, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(InsertTest(arr))
	//fmt.Println(InsertSort1(arr))
	fmt.Println(InsertSort2(arr))
}
