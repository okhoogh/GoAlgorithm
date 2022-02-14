package P02_ArrayStack

import (
	"errors"
	"fmt"
)

type ArrayStack interface {
	Clear()                    // 清空
	Size() int                 // 获取栈大小
	Push(val interface{})      // 压栈
	Pop() (interface{}, error) // 弹栈
	Top() (interface{}, error) // 获取栈顶
	IsFull() bool              // 判断栈是不是满
	IsEmpty() bool             // 判断栈是不是空
	String() string            // 返回字符串
}

type Stack struct {
	data []interface{} // 数据
	size int           // 大小
	cap  int           // 容量
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.data = make([]interface{}, 0, 10)
	stack.cap = 10
	stack.size = 0
	return stack
}

func (st *Stack) Clear() {
	st.data = make([]interface{}, 0, 10)
	st.size = 0
	st.cap = 10
}

func (st *Stack) Size() int {
	return st.size
}

func (st *Stack) autoScale() {
	newStack := make([]interface{}, 0, 2*st.cap)
	copy(newStack, st.data)
	st.cap *= 2
}

func (st *Stack) Push(val interface{}) {
	if st.IsFull() {
		st.autoScale()
	}
	st.data = append(st.data, val)
	st.size++
}

func (st *Stack) Pop() (interface{}, error) {
	if st.IsEmpty() {
		return nil, errors.New("the stack is empty")
	}
	val := st.data[st.size-1]
	st.data = st.data[:st.size-1]
	st.size--
	return val, nil
}

func (st *Stack) Top() (interface{}, error) {
	if st.IsEmpty() {
		return nil, errors.New("the stack is empty")
	}
	return st.data[st.size-1], nil
}

func (st *Stack) IsFull() bool {
	return st.size >= st.cap
}

func (st *Stack) IsEmpty() bool {
	return st.size == 0
}

func (st *Stack) String() string {
	return fmt.Sprint(st.data)
}
