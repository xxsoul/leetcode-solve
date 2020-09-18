package cn202009

import (
	"fmt"
)

// Code0917 2020-09-17答案
/*
1.遍历所有边，计算字节点的父节点，如果子节点拥有超过一个父节点，则这条边是产生冲突的边，这个子节点为产生冲突的节点。
2.如果边不是冲突的边，通过查找子节点和其父节点的祖先，如果祖先相同，则这条边是产生环路的边。如果祖先不同，则将子节点和父节点的祖先在并查集里连接。
3.一个冲突的节点，其中一个父节点一定处于一个环路里，则冲突节点在环路里的父节点的这条边，就是附加的边。
*/
func Code0917() {
	edges := [][]int{}

	edges = append(edges, []int{1, 2})
	edges = append(edges, []int{2, 3})
	edges = append(edges, []int{3, 4})
	edges = append(edges, []int{4, 1})
	edges = append(edges, []int{1, 5})

	fmt.Println(findRedundantDirectedConnection(edges))
}

func findRedundantDirectedConnection(edges [][]int) (res []int) {
	parent := make([]int, 1000) //记录节点的父节点
	conflictPath := [][]int{}   //记录冲突路径
	loopPath := [][]int{}       //纪录环路的路径

	for _, path := range edges {
		pathStart, pathEnd := path[0], path[1]

		if parent[pathStart] == 0 {
			parent[pathStart] = pathStart //父节点就是自己
		}
		if parent[pathEnd] == 0 {
			parent[pathEnd] = pathEnd //父节点就是自己
		}
		if parent[pathEnd] == pathEnd {
			//判断是否构成环路
			if findAncestry(pathEnd, parent) == findAncestry(pathStart, parent) {
				//构成回环，当前路径记录到环路路径中
				loopPath = append(loopPath, path)
				continue
			}

			parent[pathEnd] = pathStart //记录父节点
		} else if parent[pathEnd] != pathStart {
			//子节点包含多个父节点，当前路径记录到冲突路径里
			conflictPath = append(conflictPath, path)
		}
	}

	//没有构成环路，那么最后一条冲突路径就是解
	if len(loopPath) == 0 {
		return conflictPath[len(conflictPath)-1]
	}
	//没有构成冲突，那么最后一条环路路径就是解
	if len(conflictPath) == 0 {
		return loopPath[len(loopPath)-1]
	}

	//倒序查冲突路径，获取冲突节点
conflicMark:
	for i := len(conflictPath) - 1; i >= 0; i-- {
		conflictNode := conflictPath[i][1]
		//顺着环路路径往回找，遇到子节点与冲突节点一致，那这个路径就是造成环路的冲突路径，直接返回
		for _, path := range loopPath {
			curNode := path[0]
			parentNode := parent[path[0]]

			for curNode != parentNode {
				if curNode == conflictNode {
					res = []int{parentNode, curNode}
					break conflicMark
				}

				curNode = parentNode
				parentNode = parent[curNode]
			}

		}
	}

	return res
}

// findAncestry 查找给定节点的祖先
func findAncestry(node int, parent []int) (ancestry int) {
	curNode := node
	parentNode := parent[curNode]

	for curNode != parentNode {
		curNode = parentNode
		parentNode = parent[curNode]
	}
	ancestry = curNode

	return ancestry
}

// findRedundantDirectedConnectionOffical 官方解题答案
func findRedundantDirectedConnectionOffical(edges [][]int) (redundantEdge []int) {
	numNodes := len(edges)
	uf := newUnionFind(numNodes + 1)
	parent := make([]int, numNodes+1) // parent[i] 表示 i 的父节点
	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) { // from 和 to 已连接
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}

	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，其中之一与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
