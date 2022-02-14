package main

import (
	"./P03_Queue"
	"fmt"
	"io/ioutil"
)

func main411() {
	q := P03_Queue.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	fmt.Println(q.String())

	deq, _ := q.Dequeue()
	fmt.Println(deq)
	fmt.Println(q.String())

	front, _ := q.Front()
	end, _ := q.End()
	fmt.Println(front, end)

	for i := 0; i < 10; i++ {
		if ft, err := q.Dequeue(); err == nil {
			fmt.Println(ft)
		} else {
			fmt.Println(ft, err)
		}
	}
}

func main412() {
	path := "I:\\yin\\go\\10.2.go专业数据结构与算法\\1.数组队列与栈"
	files := []string{}

	q := P03_Queue.NewQueue()
	q.Enqueue(path)

	item := 0
	for !q.IsEmpty() {
		front, _ := q.Dequeue()
		dirs, _ := ioutil.ReadDir(front.(string))
		for _, dir := range dirs {
			item++
			if dir.IsDir() {
				newPath := front.(string) + "\\" + dir.Name()
				files = append(files, newPath)
				q.Enqueue(newPath)
			} else {
				newPath := front.(string) + "\\" + dir.Name()
				files = append(files, newPath)
			}
		}
	}

	fmt.Println(item)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func main() {
	var cq P03_Queue.CircularQueue
	P03_Queue.InitCircularQueue(&cq)

	for i := 1; i < 6; i++ {
		err := P03_Queue.Enqueue(&cq, i)
		if err != nil {
			fmt.Println(err)
		}
	}

	//fmt.Println(P03_Queue.Dequeue(&cq))
	//fmt.Println(P03_Queue.Dequeue(&cq))
	//fmt.Println(P03_Queue.Dequeue(&cq))
	//fmt.Println(P03_Queue.Dequeue(&cq))
	//fmt.Println(P03_Queue.Dequeue(&cq))
}
