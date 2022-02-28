package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	username int
	password string
}

func QuickSort(arr []QQ) []QQ {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	splitVal := arr[rand.Intn(length)].username
	low := make([]QQ, 0)
	high := make([]QQ, 0)
	mid := make([]QQ, 0)
	for i := 0; i < length; i++ {
		if arr[i].username < splitVal {
			low = append(low, arr[i])
		} else if arr[i].username > splitVal {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}

func BinarySearch(arr []QQ, val int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if val < arr[mid].username {
			right = mid - 1
		} else if val > arr[mid].username {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	path := "D:\\Data\\QQ.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)
	const N = 10000000
	length := 0
	usrArr := make([]QQ, N, N)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF || length == N {
			break
		}
		str := string(line)
		strArr := strings.Split(str, "----")
		if len(strArr) == 2 {
			usrArr[length].username, _ = strconv.Atoi(strArr[0])
			usrArr[length].password = strArr[1]
			length++
		}
	}
	fmt.Println("Data loading completed...")

	now := time.Now()
	usrArr = QuickSort(usrArr)
	fmt.Println("Sort total cost", time.Since(now))

	for {
		fmt.Print("Please enter the data to query: ")
		var QQ int
		fmt.Scanf("%d", &QQ) //查询QQ
		now := time.Now()
		index := BinarySearch(usrArr, QQ)
		fmt.Println("Query total cost", time.Since(now))
		if index != -1 {
			fmt.Println("search:", usrArr[index].username, usrArr[index].password)
		} else {
			fmt.Println("no data")
		}
	}
}
