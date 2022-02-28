package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	path := "D:\\data\\QQ.txt" // 文件路径
	file, _ := os.Open(path)   // 文件指针
	defer file.Close()

	br := bufio.NewReader(file)
	i := 0
	start := time.Now()
	for {
		line, _, err := br.ReadLine()
		//_, _, err := br.ReadLine()
		if err != nil {
			break
		}
		//i++
		lineStr := string(line)
		if strings.Contains(lineStr, "zhangsan") {
			//fmt.Println(lineStr)
			i++
		}
	}
	fmt.Println("time:", time.Since(start)) // 6.6s
	fmt.Println(i)
	// total: 84331445
}
