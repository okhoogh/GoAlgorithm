package main

import (
	"./P01_ArrayList"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetAll(path string, files []string) ([]string, error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.New("dir can not read")
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			fullPath := path + "\\" + dir.Name()
			files = append(files, fullPath)
			files, err = GetAll(fullPath, files)
			if err != nil {
				return nil, errors.New("dir can not read")
			}
		} else {
			fullPath := path + "\\" + dir.Name()
			files = append(files, fullPath)
		}
	}
	return files, err
}

func main111() {
	path := "I:\\yin\\go\\10.2.go专业数据结构与算法\\1.数组队列与栈"
	files := []string{}
	files, _ = GetAll(path, files)

	fmt.Println(len(files))
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func main112() {
	path := "I:\\yin\\go\\10.2.go专业数据结构与算法\\1.数组队列与栈"
	files := []string{}
	st := P01_ArrayList.NewListStack()
	st.Push(path)
	for !st.IsEmpty() {
		newPath, _ := st.Pop()
		newPath = newPath.(string)
		dirs, err := ioutil.ReadDir(newPath.(string))
		if err != nil {
			fmt.Println("dir can not read")
			break
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				fullPath := newPath.(string) + "\\" + dir.Name()
				files = append(files, fullPath) // 追加目录名
				st.Push(fullPath)
			} else {
				fullPath := newPath.(string) + "\\" + dir.Name()
				files = append(files, fullPath) // 追加文件名
			}
		}
	}

	fmt.Println(len(files))
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
