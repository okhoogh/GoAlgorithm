package main

import (
	"fmt"
	"strings"
)

func StringSelectSortOpr(strs []string) []string {
	size := len(strs)
	if size == 1 {
		return strs
	}
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if strs[min] > strs[j] {
				min = j
			}
		}
		if i != min {
			strs[i], strs[min] = strs[min], strs[i]
		}
	}
	return strs
}

func StringSelectSortCompare(strs []string) []string {
	size := len(strs)
	if size == 1 {
		return strs
	}
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if strings.Compare(strs[min], strs[j]) > 0 {
				min = j
			}
		}
		if i != min {
			strs[i], strs[min] = strs[min], strs[i]
		}
	}
	return strs
}

func main333() {
	strs := []string{"zz", "cc", "aa", "dd"}
	sort := StringSelectSortCompare(strs)
	fmt.Println(sort)
}
