package main

import "fmt"

// 1  9   2   8    3 7  4  6  5  10

// Round 1
// 1                  7
//    9                  4
//       2                 6
//            8                5
//                 3              10
// Round 2
// 1  4   2   5     3      7 9   6 8  10
// 1                3              8
//    4                 7             10
//        2                  9
//             5                 6
// Round 3
// 1  4   2   5     3      7 9   6 8  10
// 1          5               9       10
//    3            4             6
//        2               7         8
// Round 4
// 1  3  2    5     4     7   9 6  8 10
// 1      2        4         8     9
//    3       5         6       7    10
// 1  3  2   5    4     6   8   7  9  10

// ShellSortStep 根据初始点和间隔进行看似跳跃的插入排序
func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap {
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	// 1.定义初始间隔
	gap := length / 2
	for gap > 0 {
		// 2.以i为初始起点和gap为间隔遍历
		for i := 0; i < gap; i++ {
			ShellSortStep(arr, i, gap)
		}
		gap /= 2
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(ShellSort(arr))
}
