package main

import "fmt"

// 鸡尾酒排序-双向冒泡排序
// 形象比喻：来回倒
// 第一轮左半边数组最大值移到右半边数组第一个位置
// 因为第一轮最后一次比较，左最大移到右最小，右最小移到左最大

func CockTailSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length/2; i++ {
		left, right := 0, length-1
		for left < right {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
		//fmt.Println(arr)
	}
	return arr
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("arr:", arr)
	fmt.Println("sorted:", CockTailSort(arr))
}
