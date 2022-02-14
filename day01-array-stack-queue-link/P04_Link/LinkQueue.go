package P04_Link

import (
	"errors"
)

type LinkQueue struct {
	head *Node
	tail *Node
}

type QueueLink interface {
	Length() int
	EnQueue(val interface{})
	DeQueue() (interface{}, error)
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{}
}

func (lq *LinkQueue) Length() int {
	length := 0
	for n := lq.head; n != nil; n = n.Next {
		length++
	}
	return length
}

func (lq *LinkQueue) EnQueue(val interface{}) {
	newNode := &Node{val, nil}
	if lq.head == nil {
		lq.head = newNode
		lq.tail = newNode
	} else {
		lq.tail.Next = newNode
		lq.tail = newNode
	}
}

func (lq *LinkQueue) DeQueue() (interface{}, error) {
	if lq.head == nil {
		return nil, errors.New("linkQueue is empty")
	}
	val := lq.head.Data
	if lq.head == lq.tail {
		lq.head, lq.tail = nil, nil
		return val, nil
	}
	lq.head = lq.head.Next
	return val, nil
}
