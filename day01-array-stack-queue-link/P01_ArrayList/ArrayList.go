package P01_ArrayList

import (
	"errors"
	"fmt"
)

// List 接口
type List interface {
	Size() int                               // 获取数组大小
	Get(index int) (interface{}, error)      // 获取第几个数
	Set(index int, val interface{}) error    // 修改数组
	Insert(index int, val interface{}) error // 插入元素
	Append(val interface{})                  // 追加元素
	Delete(index int) error                  // 删除元素
	Clear()                                  // 清空数组
	String() string                          // 返回数组的字符串
	Iterator() Iterator                      // 返回迭代器
}

// ArrayList 数据结构
type ArrayList struct {
	data []interface{}
	size int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.data = make([]interface{}, 0, 10)
	list.size = 0
	return list
}

func (list *ArrayList) checkFull() {
	if list.size == cap(list.data) {
		newData := make([]interface{}, 2*list.size)
		copy(newData, list.data)
		list.data = newData
	}
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index out of bound")
	}
	return list.data[index], nil
}

func (list *ArrayList) Append(val interface{}) {
	list.data = append(list.data, val)
	list.size++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.data)
}

func (list *ArrayList) Set(index int, val interface{}) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bound")
	}
	list.data[index] = val
	return nil
}

func (list *ArrayList) Insert(index int, val interface{}) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bound")
	}
	list.checkFull()
	list.data = list.data[:list.size+1]
	for i := list.size; i > index; i-- {
		list.data[i] = list.data[i-1]
	}
	list.data[index] = val
	list.size++
	return nil
}

func (list *ArrayList) Delete(index int) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bound")
	}
	list.data = append(list.data[:index], list.data[index+1:]...)
	list.size--
	return nil
}

func (list *ArrayList) Clear() {
	list.data = make([]interface{}, 0, 10)
	list.size = 0
}
