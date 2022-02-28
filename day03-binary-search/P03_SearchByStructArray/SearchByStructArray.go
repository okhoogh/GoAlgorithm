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

type QQ struct {
	username int
	password string
}

func main() {
	path := "D:\\data\\QQ.txt"
	file, _ := os.Open(path)
	defer file.Close()

	br := bufio.NewReader(file)
	const N = 84331445
	length := 0
	usrArr := make([]QQ, N, N)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
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

	for {
		fmt.Print("Please enter the data to query: ")
		var QQ int
		fmt.Scanf("%d", &QQ) //查询QQ
		now := time.Now()
		for i := 0; i < N; i++ {
			if QQ == usrArr[i].username {
				fmt.Println(usrArr[i].username, usrArr[i].password)
				fmt.Println("The total cost", time.Since(now))
			}
		}
	}
}
