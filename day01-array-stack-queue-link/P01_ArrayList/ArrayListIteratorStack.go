package P01_ArrayList

import (
	"errors"
)

type ListIteratorStack interface {
	Clear()                    // 清空
	Size() int                 // 获取栈大小
	Push(val interface{})      // 压栈
	Pop() (interface{}, error) // 弹栈
	Top() (interface{}, error) // 获取栈顶
	IsFull() bool              // 判断栈是不是满
	IsEmpty() bool             // 判断栈是不是空
	String() string            // 返回字符串
}

type ArrayListIteratorStack struct {
	data     *ArrayList // 数据
	cap      int        // 容量
	Iterator Iterator
}

func NewArrayListIteratorStack() *ArrayListIteratorStack {
	stack := new(ArrayListIteratorStack)
	stack.data = NewArrayList()
	stack.cap = 10
	stack.Iterator = stack.data.Iterator()
	return stack
}

func (st *ArrayListIteratorStack) Clear() {
	st.data = NewArrayList()
	st.cap = 10
}

func (st *ArrayListIteratorStack) Size() int {
	return st.data.Size()
}

func (st *ArrayListIteratorStack) Push(val interface{}) {
	st.data.Append(val)
}

func (st *ArrayListIteratorStack) Pop() (interface{}, error) {
	if st.IsEmpty() {
		return nil, errors.New("the stack is empty")
	}
	val, _ := st.data.Get(st.data.Size() - 1)
	st.data.Delete(st.data.Size() - 1)
	return val, nil
}

func (st *ArrayListIteratorStack) Top() (interface{}, error) {
	if st.IsEmpty() {
		return nil, errors.New("the stack is empty")
	}
	top, err := st.data.Get(st.data.Size() - 1)
	return top, err
}

func (st *ArrayListIteratorStack) IsFull() bool {
	return st.data.Size() >= st.cap
}

func (st *ArrayListIteratorStack) IsEmpty() bool {
	return st.data.Size() == 0
}

func (st *ArrayListIteratorStack) String() string {
	return st.data.String()
}
