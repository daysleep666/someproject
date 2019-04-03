package main

import "fmt"

// 带权有向图

const (
	max = 99999999
)

type Graph struct {
	E map[int][]*Edge
}

type Edge struct {
	To     int
	Weight int
}

func (g *Graph) Add(f, t, w int) {
	g.E[f] = append(g.E[f], &Edge{To: t, Weight: w})
	g.E[t] = append(g.E[t], &Edge{To: f, Weight: w})
}

func (g *Graph) dijkstra(node int) {
	u, s := make(map[int]int), make(map[int]int)
	// 初始化当前节点的距离为0，其他节点的距离为max
	for k, _ := range g.E {
		u[k] = max
	}
	u[node] = 0
	s[node] = 0

	list := []int{node}
	for len(list) != 0 {
		for i := 0; i < len(list); i++ {
			cur := list[i]
			// 1.找出当前节点到所有其他节点的距离更新到u中
			es := g.E[cur]
			for _, e := range es {
				// 如果从初始节点到当前节点的距离加上当前节点到新节点的距离小于u中记录的距离
				newWeight := (s[cur] + e.Weight)
				if newWeight < u[e.To] {
					u[e.To] = newWeight
				}
			}
			// 2.选出距离最短的节点加入s中
			minEdges := make([]*Edge, 0)
			for k, v := range u {
				_, isExist := s[k]
				if (len(minEdges) == 0 || minEdges[0].Weight >= v) && !isExist {
					if len(minEdges) == 0 || minEdges[0].Weight > v {
						minEdges = []*Edge{&Edge{To: k, Weight: v}}
					} else {
						minEdges = append(minEdges, &Edge{To: k, Weight: v})
					}
				}
			}
			list = make([]int, 0, len(minEdges))
			for _, v := range minEdges {
				s[v.To] = v.Weight
				list = append(list, v.To)
			}
			// 重复1
		}
	}

	for k, v := range s {
		fmt.Printf("%v->%v:%v \n", node, k, v)
	}

}

func main() {
	g := Graph{make(map[int][]*Edge, 0)}
	g.Add(1, 2, 1)
	g.Add(1, 3, 3)
	g.Add(1, 4, 4)
	g.Add(2, 3, 1)
	g.Add(3, 4, 1)
	g.Add(2, 5, 1)
	// for k, v := range g.E {
	// 	for _, vv := range v {
	// 		fmt.Printf("%v->%v:%v\n", k, vv.To, vv.Weight)
	// 	}
	// }
	g.dijkstra(1)
}
