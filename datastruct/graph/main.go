package main

import (
	"fmt"
)

// 树的特效 一对多
// 图的特性 多对多

type DataNode struct {
	Data     int
	NextNode *DataNode
}

func (n *DataNode) Add(_data int) {
	var tmpNode = n
	for tmpNode.NextNode != nil {
		tmpNode = tmpNode.NextNode
	}
	tmpNode.NextNode = &DataNode{
		Data: _data,
	}
}

func (n *DataNode) Display() {
	var tmpNode = n
	for tmpNode != nil {
		fmt.Printf("->%v", tmpNode.Data)
		tmpNode = tmpNode.NextNode
	}
}

type DataGraph []*DataNode // 邻接矩阵

func Init(_max int) DataGraph {
	dg := make([]*DataNode, _max, _max)
	return dg
}

func (dg *DataGraph) Add(_who, _inviteWho int) error {
	if _who >= dg.Len() { // 重新分配足够大的内存
		newDG := Init(_who + 1)
		copy(newDG, *dg)
		*dg = newDG //
	}
	if (*dg)[_who] == nil {
		(*dg)[_who] = &DataNode{Data: _who}
	}
	(*dg)[_who].Add(_inviteWho)
	return nil
}

func (dg *DataGraph) Display() {
	for _, v := range *dg {
		if v != nil {
			v.Display()
			fmt.Println()
		}
	}
}

func (dg *DataGraph) Len() int {
	return len(*dg)
}

func main() {
	dg := Init(3)
	dg.Add(1, 10)
	dg.Add(2, 3)
	dg.Add(1, 4)
	dg.Add(3, 6)
	dg.Add(4, 5)
	dg.Display()
}

// 邀请码用图来处理
/*
用邻接矩阵记录谁邀请了谁
a邀请了b，b邀请了c，c邀请了d
a->b->c->d->...

用map来记录谁被谁邀请了
m[我]=邀请我的人
这样循环n次就可以得到我的上n级到我的邀请链


*/
