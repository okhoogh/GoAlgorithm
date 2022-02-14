package main

import (
	"./P01_ArrayList"
	"./P02_ArrayStack"
	"fmt"
)

// 1.可变长数组追加整型
func main1() {
	list := P01_ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list.String())
}

// 2.可变长数组追加字符串
func main2() {
	list := P01_ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list.String())
}

// 3.未实现完所有接口的类赋值不了给对应接口
/*
func main3() {
	var list P01_ArrayList.List = P01_ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list.String())
}
*/

// 4.实现完所有接口的类可以赋值给对应接口
func main4() {
	var list P01_ArrayList.List = P01_ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list.String())
}

// 5.测试可变长数组的所有接口
func main5() {
	var list P01_ArrayList.List = P01_ArrayList.NewArrayList()

	fmt.Println("---append---")
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Insert(1, "d")
	for i := 0; i < 10; i++ {
		list.Insert(1, "d")
	}
	fmt.Println(list.String())

	fmt.Println("---delete---")
	list.Delete(0)
	fmt.Println(list.String())
}

// 6.测试可变长数组的所有接口
func main6() {
	var list P01_ArrayList.List = P01_ArrayList.NewArrayList()

	fmt.Println("---append---")
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Insert(1, "d")
	for i := 0; i < 10; i++ {
		list.Insert(1, "d")
	}
	fmt.Println(list.String())

	fmt.Println("---delete---")
	list.Delete(0)
	fmt.Println(list.String())
}

// 7.测试可变长数组的迭代器
func main7() {
	var list P01_ArrayList.List = P01_ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	list.Append("d4")
	list.Append("e5")

	for it := list.Iterator(); it.HasNext(); {
		next, _ := it.Next()
		if next == "c3" {
			it.Remove()
		}
		fmt.Println(next)
	}
	fmt.Println(list)
}

// 8.测试栈
func main8() {
	st := P02_ArrayStack.NewStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	fmt.Println(st)
	v1, _ := st.Pop()
	fmt.Println(v1, st.String())

	top, _ := st.Top()
	fmt.Println(top, st.IsEmpty(), st.IsFull())

	st.Clear()
	fmt.Println(st.IsEmpty())
}

// 9.测试数组栈
func main9() {
	st := P01_ArrayList.NewListStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	fmt.Println(st)
	v1, _ := st.Pop()
	fmt.Println(v1, st.String())

	top, _ := st.Top()
	fmt.Println(top, st.IsEmpty(), st.IsFull())

	st.Clear()
	fmt.Println(st.IsEmpty())
}

// 10.测试有迭代器数组栈
func main10() {
	st := P01_ArrayList.NewArrayListIteratorStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	for it := st.Iterator; it.HasNext(); {
		next, _ := it.Next()
		fmt.Println(next)
	}
}

// Add 递归求和
func Add(num int) int {
	if num == 0 {
		return 0
	}
	return num + Add(num-1)
}

// 11.递归求和转成循环+手写栈
func main11() {
	//fmt.Println(Add(5))

	st := P02_ArrayStack.NewStack()
	st.Push(5)
	sum := 0
	for !st.IsEmpty() {
		val, err := st.Pop()
		if err != nil || val == 0 {
			break
		}
		sum += val.(int)
		st.Push(val.(int) - 1)
	}
	fmt.Println(sum)
}

// Fibonacci 递归斐波那契
func Fibonacci(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}

// 递归斐波那契转成循环+手写栈
func main() {
	//fmt.Println(Fibonacci(6))
	st := P02_ArrayStack.NewStack()
	st.Push(6)
	num := 0
	for !st.IsEmpty() {
		val, err := st.Pop()
		if err != nil {
			break
		}
		if val == 1 || val == 2 {
			num += 1
		} else {
			st.Push(val.(int) - 1)
			st.Push(val.(int) - 2)
		}
	}
	fmt.Println(num)
}
