package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("D:\\Data\\CSDN600w.sql")
	if err != nil {
		fmt.Println("file read failed")
		return
	}
	defer file.Close()

	savePath := "D:\\CSDN-User.txt"
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
		fmt.Fprintln(save, lineStrArr[0])
	}
	save.Flush()
}
