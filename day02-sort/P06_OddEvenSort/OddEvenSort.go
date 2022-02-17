package main

import "fmt"

func OddEvenSort(arr []int) []int {
	// 是否有序标志
	isSorted := false
	length := len(arr)
	for !isSorted {
		// 假设有序
		isSorted = true
		// 排偶数（从奇数位开始跟偶数位进行比较交换）
		for i := 0; i < length-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false // 如果交换过，说明数组还未有序
			}
		}
		// 排奇数（从偶数位开始跟奇数位进行比较交换）
		for i := 1; i < length-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(OddEvenSort(arr))
}
