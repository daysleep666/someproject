package graph

import "fmt"

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
