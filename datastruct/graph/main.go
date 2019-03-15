package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/graph/graph"
)

func main() {
	dg := graph.NewGraph()
	dg.Push(1, 2)
	dg.Push(1, 3)
	dg.Push(1, 4)
	dg.Push(2, 3)
	dg.Push(4, 5)

	dg.BFS(1)
	fmt.Println(dg.M)

}

// 树的特效 一对多
// 图的特性 多对多
// 邀请码用图来处理
/*
用邻接矩阵记录谁邀请了谁
a邀请了b，b邀请了c，c邀请了d
a->b->c->d->...

用map来记录谁被谁邀请了
m[我]=邀请我的人
这样循环n次就可以得到我的上n级到我的邀请链


*/
