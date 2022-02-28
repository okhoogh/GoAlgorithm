package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func QuickSortRandom(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	splitNum := arr[rand.Intn(length)]
	low := make([]string, 0)
	high := make([]string, 0)
	mid := make([]string, 0)
	for i := 0; i < length; i++ {
		if splitNum > arr[i] {
			low = append(low, arr[i])
		} else if splitNum < arr[i] {
			high = append(high, arr[i])
		} else {
			mid = append(mid, splitNum)
		}
	}
	low, high = QuickSortRandom(low), QuickSortRandom(high)
	return append(append(low, mid...), high...)
}

func TestLoop() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
	}
	duration := time.Since(start)
	fmt.Println(duration)
}

func TestReadWrite() {
	start := time.Now()
	file, err := os.Open("D:\\Data\\CSDN600w.sql")
	if err != nil {
		fmt.Println("file read failed")
		return
	}
	defer file.Close()

	savePath := "D:\\Data\\CSDN-Mail.txt"
	saveFile, _ := os.Create(savePath)
	defer saveFile.Close()

	save := bufio.NewWriter(saveFile)
	br := bufio.NewReader(file)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		lineStrArr := strings.Split(lineStr, " # ")
		//fmt.Println(lineStrArr[1])
		//save.WriteString(lineStrArr[1])
		fmt.Fprintln(save, lineStrArr[2])
	}
	save.Flush()
	duration := time.Since(start)
	fmt.Println(duration)
}

func TestSaveMemory() {
	start := time.Now()
	const N = 6428632
	strs := make([]string, N, N)

	file, err := os.Open("D:\\Data\\CSDN-Mail.txt")
	if err != nil {
		fmt.Println("file read failed")
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	i := 0
	for {
		mail, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		strs[i] = string(mail)
		i++
		//strs = append(strs, string(mail))
	}
	//time.Sleep(time.Second * 100)
	duration := time.Since(start)
	fmt.Println(duration)
}

func main() {
	const N = 6428632
	strs := make([]string, N, N)

	file, err := os.Open("D:\\Data\\CSDN-Pass.txt")
	if err != nil {
		fmt.Println("file read failed")
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for i := 0; ; i++ {
		mail, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		strs[i] = string(mail)
	}
	fmt.Println("read file complete")

	start := time.Now()
	strs = QuickSortRandom(strs)
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println("sort complete")

	savePath := "D:\\Data\\CSDN-PassSort.txt"
	saveFile, _ := os.Create(savePath)
	defer saveFile.Close()
	save := bufio.NewWriter(saveFile)
	for i := 0; i < len(strs); i++ {
		fmt.Fprintln(save, strs[i])
	}
	save.Flush()
}
