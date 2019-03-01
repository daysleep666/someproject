package main

import "github.com/daysleep666/someproject/datastruct/graph/graph"

func main() {
	dg := graph.Init()
	dg.Add("0", "1")
	dg.Add("0", "3")
	dg.Add("1", "2")
	dg.Add("2", "4")
	// dg.Add("3", "4")

	dg.BFS("0")
	dg.DFS("0")
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
