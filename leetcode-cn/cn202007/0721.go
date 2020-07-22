package cn202007

import "fmt"

func Code0721() {
	aaa := helper(1, 5)

	fmt.Printf(string(len(aaa)))
}

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{nil}
	}

	// 构建一个切片，枚举可能给定的值
	numArr := make([]*TreeNode, n)
	for i := 0; i <= n; i++ {
		numArr[i] = &TreeNode{Val: i + 1}
	}

	reset := func(rootIdx int) {
		for index, node := range numArr {
			if index == 0 {
				continue
			}
			if index <= rootIdx {
				node.Left = numArr[index-1]
			}
			if index > rootIdx {
				numArr[index-1].Right = node
			}
		}
	}

	result := make([]*TreeNode, 0)
	// 需要多层循环
	// 第一层循环，用于最外层根节点的值和在数组上的索引
	for globalRootIdx, _ := range numArr {
		// 这个时候需要恢复整棵二叉树为最简单状态，也就是以根为单位，左子树倒序，右子树升序
		reset(globalRootIdx)
		// rootNode := globalRootNode

		// 遍历左子树
	}

	return result
}

// helper 官方答案
func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}

// checkNodeCanTilt 验证节点是否可以向给定方向倾斜
// direct 0-left;1-right
func checkNodeCanTilt(node *TreeNode, direct int) bool {
	if direct == 0 && node.Right == nil || direct > 0 && node.Left == nil {
		return false
	}
	return true
}

// nodeTilt 节点向给定方向倾斜（让给定节点的左子树或右子树取代当前位置）
// direct 0-left;1-right
func nodeTilt(node *TreeNode, parent *TreeNode, direct int) {
	if direct == 0 {
		// 向左倾斜
		// 判断当前是父节点的左孩子还是右孩子
		if parent.Left == node {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
		node.Right.Left = node
		node.Right = nil
	} else {
		//向右倾斜
		if parent.Left == node {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
		node.Left.Right = node
		node.Left = nil
	}
}

// linkNode 连接两个节点
// par: 父节点
// chi: 子节点
// return: 子节点挂到了父节点哪个指针上，0-left；1-right
func linkNode(par *TreeNode, chi *TreeNode) int {
	if par.Val <= chi.Val {
		par.Right = chi
		return 1
	}
	par.Left = chi
	return 0
}

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
