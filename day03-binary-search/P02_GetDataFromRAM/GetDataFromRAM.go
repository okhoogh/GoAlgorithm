package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	path := "D:\\data\\QQ.txt" // 文件路径
	file, _ := os.Open(path)   // 文件指针
	defer file.Close()         // 延迟

	br := bufio.NewReader(file)
	length := 0
	const N = 84331445
	strArr := make([]string, N, N)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		strArr[length] = string(line)
		length++
	}
	fmt.Println("数据加载完成")

	for {
		cnt := 0
		searchStr := ""
		fmt.Print("请输入所要查找的字符串: ")
		fmt.Scan(&searchStr)
		start := time.Now()
		for i := 0; i < N; i++ {
			if strings.Contains(strArr[i], searchStr) {
				//fmt.Println(lineStr)
				cnt++
			}
		}
		fmt.Println("time:", time.Since(start))
		fmt.Println(cnt)
	}
}
