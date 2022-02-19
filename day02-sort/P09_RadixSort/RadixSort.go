package main

import (
	"errors"
	"fmt"
)

//银行，每人都有存款
//100 1000  10000
//99, 799, 123, 246, 1650, 2370
//银行存款，金额固定，不固定

//800万高考考生  分数排序 0-750 1500
// 687 3 688 6 689  8 690

//10亿人，身高排序   0-300

/**
https://www.cnblogs.com/sun/archive/2008/06/26/1230095.html
算法步骤：
1.先找出数组中最大值
2.根据最大值的十进制位数进行分段排序(个位、十位、百位、千位...)
3.定义一个长度为10的位数数组，用于存储原数组的所有数在每一位(个位、十位、百位、千位...)的个数
4.将位数数组从前到后进行累加，用于将原数组中的数映射到根据位数划分的子数组中
位数：[0  1  2   3   4   5   6  7    8   9]
位数组[0  6  4   6   0   0   0  0    4   0]
累加中[0 +0 +6 +10 +16 +16 +16 +16 +16 +20]
新数组[0  6 10  16  16  16  16  16  20  20]
意思：位数为0的数有0个，不占坑
	 位数为1的数有6个，占6个坑，于是从0开始占6个坑，[0, 6)都是(个位、十位、百位、千位...)为1的数
	 位数为2的数有4个，占4个坑，于是从6开始占4个坑，[6, 10)都是(个位、十位、百位、千位...)为2的数
	 位数为3的数有6个，占6个坑，于是从10开始占6个坑，[10, 16)都是(个位、十位、百位、千位...)为3的数
	 位数为4的数有0个，不占坑
	 位数为5的数有0个，不占坑
	 位数为6的数有0个，不占坑
	 位数为7的数有0个，不占坑
	 位数为8的数有4个，占4个坑，于是从16开始占4个坑，[16, 20)都是(个位、十位、百位、千位...)为8的数
	 位数为9的数有0个，不占坑
*/

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

func RadixSort(arr []int) []int {
	max, _ := SelectGetMax(arr)
	for bit := 1; max/bit > 0; bit *= 10 {
		//fmt.Println(bit)
		arr = SortByBit(arr, bit)
	}
	fmt.Println("-----------------")
	return arr
}

func SortByBit(arr []int, bit int) []int {
	length := len(arr)
	bitCount := make([]int, 10)
	for i := 0; i < length; i++ {
		num := arr[i] / bit % 10
		bitCount[num]++
	}
	fmt.Println("bitCount1:", bitCount)
	for i := 1; i < 10; i++ {
		bitCount[i] += bitCount[i-1]
	}
	fmt.Println("bitCount2:", bitCount)

	tmp := make([]int, length)
	//for i := 0; i < length; i++ {
	for i := length - 1; i >= 0; i-- {
		num := arr[i] / bit % 10
		tmp[bitCount[num]-1] = arr[i]
		bitCount[num]--
	}
	fmt.Println("tmp:", tmp)
	return tmp
}

func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 111011, 11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 111011}
	fmt.Println("arr:", arr)
	fmt.Println("----------------------------")
	fmt.Println("sorted:", RadixSort(arr))
}
