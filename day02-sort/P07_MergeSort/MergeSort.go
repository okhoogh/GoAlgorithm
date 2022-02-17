package main

import "fmt"

func merge(left, right []int) []int {
	l, r := 0, 0
	llen, rlen := len(left), len(right)
	tmp := make([]int, 0)
	for l < llen && r < rlen {
		if left[l] < right[r] {
			tmp = append(tmp, left[l])
			l++
		} else if left[l] > right[r] {
			tmp = append(tmp, right[r])
			r++
		} else {
			tmp = append(tmp, left[l], right[r])
			l++
			r++
		}
	}
	if l < llen {
		tmp = append(tmp, left[l:]...)
	}
	if r < rlen {
		tmp = append(tmp, right[r:]...)
	}
	return tmp
}

func MergeSort(arr []int) []int {
	length := len(arr)
	// todo: 小于10改用插入排序
	if length <= 1 {
		return arr
	}
	mid := length / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(MergeSort(arr))
}
