package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	path := "D:\\data\\QQ.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)
	const N = 84331445
	length := 0
	usrArr := make(map[int]string, N)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(line)
		strArr := strings.Split(str, "----")
		if len(strArr) == 2 {
			username, _ := strconv.Atoi(strArr[0])
			password := strArr[1]
			usrArr[username] = password
			length++
		}
	}
	fmt.Println("Data loading completed...")

	for {
		fmt.Print("Please enter the data to query: ")
		var QQ int
		fmt.Scanf("%d", &QQ) //查询QQ
		now := time.Now()
		for i := 0; i < N; i++ {
			password, err := usrArr[i]
			if err {
				fmt.Println(QQ, password)
				fmt.Println("The total cost: ", time.Since(now))
			} else {
				fmt.Println("No data")
			}
		}
	}
}
