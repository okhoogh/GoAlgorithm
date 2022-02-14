package P03_Queue

import (
	"errors"
	"fmt"
)

type Queue interface {
	Size() int                     // 获取长度
	IsEmpty() bool                 // 判断队列是否为空
	isFull() bool                  // 判断队列是否已满
	Enqueue(val interface{})       // 入队
	Dequeue() (interface{}, error) // 队首出队
	Front() (interface{}, error)   // 获取队首
	End() (interface{}, error)     // 获取队尾
	String() string                // 获取字符串
	Clear()                        // 清空队列
}

type ListQueue struct {
	data []interface{}
	size int
	cap  int
}

func NewQueue() *ListQueue {
	queue := new(ListQueue)
	queue.data = make([]interface{}, 0)
	queue.size = 0
	queue.cap = 10
	return queue
}

func (q *ListQueue) Size() int {
	return q.size
}

func (q *ListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ListQueue) IsFull() bool {
	return q.size == q.cap
}

func (q *ListQueue) Enqueue(val interface{}) {
	q.data = append(q.data, val)
	q.size++
}

func (q *ListQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	front := q.data[0]
	q.data = q.data[1:q.size]
	q.size--
	return front, nil
}

func (q *ListQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.data[0], nil
}

func (q *ListQueue) End() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.data[q.size-1], nil
}

func (q *ListQueue) Clear() {
	q.data = make([]interface{}, 10)
	q.size = 0
}

func (q *ListQueue) String() string {
	return fmt.Sprint(q.data)
}
