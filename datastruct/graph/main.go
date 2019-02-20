package main

import "github.com/daysleep666/someproject/datastruct/graph/graph"

func main() {
	dg := graph.Init()
	dg.Add("a", "b")
	dg.Add("a", "c")
	dg.Add("b", "c")
	dg.Display()
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
