package graph

import (
	"fmt"
)

// 邻接表

type Node struct {
	Key string
}

type DataGraph struct { // 邻接矩阵
	Edge map[string][]string
}

func Init() *DataGraph {
	return &DataGraph{
		Edge: make(map[string][]string),
	}
}

func (dg *DataGraph) Add(_node1, _node2 string) {
	dg.Edge[_node1] = append(dg.Edge[_node1], _node2)
	dg.Edge[_node2] = append(dg.Edge[_node2], _node1)
}

func (dg *DataGraph) Display() {
	for key, edge := range dg.Edge {
		fmt.Printf("%v", key)
		for _, v := range edge {
			fmt.Printf("->%v", v)
		}
		fmt.Println()
	}
}

func (dg *DataGraph) GetNodes(_v string) []string {
	return dg.Edge[_v]
}

func (dg *DataGraph) BFS(_from string) { // 从某个节点开始广度优先遍历
	var visisted = make(map[string]bool)
	var list = []string{_from}
	visisted[_from] = true
	fmt.Printf("%v", _from)
	var tmpList []string
	for len(list) != 0 {
		for _, v := range list {
			nodes := dg.GetNodes(v)
			for _, n := range nodes {
				if !visisted[n] {
					fmt.Printf("->%v", n)
					tmpList = append(tmpList, n)
					visisted[n] = true
				}
			}
		}
		list = tmpList
		tmpList = tmpList[:0]
	}
	fmt.Println()
}

func (dg *DataGraph) DFS(_from string) { // 从某个节点开始深度优先遍历
	var visisted = make(map[string]bool)
	visisted[_from] = true
	fmt.Printf("%v", _from)
	dg.dfs(_from, visisted)
}

func (dg *DataGraph) dfs(_from string, _visisted map[string]bool) {
	for _, v := range dg.GetNodes(_from) {
		if !_visisted[v] {
			fmt.Printf("->%v", v)
			_visisted[v] = true
			dg.dfs(v, _visisted)
		}
	}
}
