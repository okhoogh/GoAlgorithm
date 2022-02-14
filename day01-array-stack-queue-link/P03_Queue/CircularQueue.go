package P03_Queue

import "errors"

const Queuesize = 3

type CircularQueue struct {
	data [Queuesize]interface{} // 数据
	head int                    // 头指针
	tail int                    // 尾指针
}

func InitCircularQueue(q *CircularQueue) {
	q.head = 0
	q.tail = 0
}

func (q *CircularQueue) Size() int {
	return (q.tail - q.head + Queuesize) % Queuesize
}

func Enqueue(q *CircularQueue, val interface{}) error {
	newTail := (q.tail + 1) % Queuesize
	if newTail == q.head {
		return errors.New("CircularQueue is full")
	}
	q.data[newTail] = val
	q.tail = newTail
	return nil
}

func Dequeue(q *CircularQueue) (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("CircularQueue is empty")
	}
	q.head = (q.head + 1) % Queuesize
	head := q.data[q.head]
	return head, nil
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}
