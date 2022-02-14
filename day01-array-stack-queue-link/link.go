package main

import (
	"./P04_Link"
	"fmt"
)

func main511() {
	node1 := new(P04_Link.Node)
	node2 := new(P04_Link.Node)
	node3 := new(P04_Link.Node)
	node4 := new(P04_Link.Node)
	node5 := new(P04_Link.Node)

	node1.Data = 1
	node1.Next = node2
	node2.Data = 2
	node2.Next = node3
	node3.Data = 3
	node3.Next = node4
	node4.Data = 4
	node4.Next = node5
	node5.Data = 5
	node5.Next = nil

	fmt.Println(node1.Data)
	fmt.Println(node2.Data)
	fmt.Println(node3.Data)
	fmt.Println(node4.Data)
	fmt.Println(node5.Data)

	fmt.Println("----------------")

	fmt.Println(node1.Next.Next.Next.Next.Data)
}

func main512() {
	lk := P04_Link.NewLinkStack()

	for i := 0; i < 100000; i++ {
		lk.Push(i)
	}

	for data, _ := lk.Pop(); data != nil; data, _ = lk.Pop() {
		fmt.Println(data)
	}
}

func main() {
	lq := P04_Link.NewLinkQueue()

	for i := 0; i < 100000; i++ {
		lq.EnQueue(i)
	}

	for data, _ := lq.DeQueue(); data != nil; data, _ = lq.DeQueue() {
		fmt.Println(data)
	}
}
