package P01_ArrayList

import "errors"

type Iterator interface {
	HasNext() bool              // 是否存在下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
	GetIndex() int              // 获取索引
}

type Iterable interface {
	Iterator() Iterator // 构建初始化接口
}

type ArraylistIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArraylistIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArraylistIterator) HasNext() bool {
	return it.currentIndex < it.list.Size()
}

func (it *ArraylistIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("no more items")
	}
	val, err := it.list.Get(it.currentIndex)
	if err == nil {
		it.currentIndex++
	}
	return val, err
}

func (it *ArraylistIterator) Remove() {
	it.currentIndex--
	if it.list.Delete(it.currentIndex) != nil {
		it.currentIndex++
	}
}

func (it *ArraylistIterator) GetIndex() int {
	return it.currentIndex
}
