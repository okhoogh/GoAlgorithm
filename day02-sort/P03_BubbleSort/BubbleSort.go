package main

import "fmt"

func BubbleSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		flag := true
		for j := 0; j < size-1-i; j++ {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			flag = false
		}
		if flag {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(BubbleSort(arr))
}
