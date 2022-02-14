package P04_Link

import "errors"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(val interface{})
	Pop() (interface{}, error)
	Length() int
}

func NewLinkStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	return n.Next == nil
}

// Push 头插法 头部插入头部删除 跟栈类似
func (n *Node) Push(val interface{}) {
	newNode := &Node{val, nil}
	newNode.Next = n.Next
	n.Next = newNode
}

// Pop 移除头节点的下一个节点，并与下下一个节点链接
func (n *Node) Pop() (interface{}, error) {
	if n.IsEmpty() {
		return nil, errors.New("linkStack is empty")
	}
	val := n.Next.Data
	n.Next = n.Next.Next
	return val, nil
}

func (n *Node) Length() int {
	length := 0
	for ; n != nil; n = n.Next {
		length++
	}
	return length
}
