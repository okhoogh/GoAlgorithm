package main

import (
	"fmt"
	"math"
)

type Node struct {
	val     int  // 叶子数据
	checked bool // 已排序或者说无穷大 false为无穷大 注意：Node数组的checked默认值为false，所以得将无穷大定义为false
	rank    int  // 叶子排序
}

// pow x^y
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func compareAndUp(tree *[]Node, left int) {
	right := left + 1
	mid := (left - 1) / 2
	// 父亲节点存儿子节点中最小的节点
	// 除非左节点无效或者右节点有效并且比左节点大
	if !(*tree)[left].checked || ((*tree)[right].checked &&
		((*tree)[left].val > (*tree)[right].val)) {
		(*tree)[mid] = (*tree)[right]
	} else {
		// 否则就选左节点
		(*tree)[mid] = (*tree)[left]
	}
}

func TournamentSort(arr []int) []int {
	length := len(arr)
	level := 0                    // 树的层数
	res := make([]int, 0, length) // 结果
	// 求出覆盖所有元素的树的层数
	for pow(2, level) < length {
		level++
	}

	leaf := pow(2, level) // 叶子数量

	// 注意：Node数组的checked默认值为false，所以得将无穷大定义为false
	tree := make([]Node, leaf*2-1) // 树的所有节点数量
	// 用数组元素填充叶子节点
	for i := 0; i < length; i++ {
		// 叶子节点从叶子个数出发，因为叶子节点个数等于非叶子节点个数之和
		tree[leaf+i-1] = Node{arr[i], true, i}
	}

	// 从叶子层向上处理
	for i := level; i > 0; i-- {
		nodeCount := pow(2, i)
		for j := 0; j < nodeCount/2; j++ {
			left := nodeCount - 1 + j*2
			compareAndUp(&tree, left)
			//fmt.Println(tree)
		}
	}

	// 保存最顶端的数
	res = append(res, tree[0].val)
	for i := 0; i < len(arr)-1; i++ {
		winNode := tree[0].rank + leaf - 1
		tree[winNode].checked = false // 将最终取胜的节点设为无穷大
		for j := 0; j < level; j++ {
			left := winNode
			if left%2 == 0 {
				left = winNode - 1
			}
			compareAndUp(&tree, left)
			winNode = (left - 1) / 2
		}
		res = append(res, tree[0].val)
	}

	return res
}

func main() {
	// 构造随机数组作为输入
	length := 10
	hash := make(map[int]int, length)
	input := make([]int, 0)
	for i := 0; i < length; i++ {
		hash[i] = i
	}
	for k := range hash {
		input = append(input, k)
	}
	fmt.Println("arr:", input)
	fmt.Println("sorted:", TournamentSort(input))
}
