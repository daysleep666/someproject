package main

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

func (g *Graph) dijkstra() {
	s, u := make(map[int]int), make(map[int]int)
	for _, e := range g.E {
		for k, _ := range e {
			u[k] = max
		}
	}
	first := 1
	s[first] = 0
	list := []int{first}
	for len(list) != 0 {
		for _, v := range list {
			nodes := g.E[v]
			for _, node := range nodes {
				u[node.To] = node.Weight
			}
		}
	}

	mink, minv := 0, max
	for k, v := range u {
		if v < minv {
			mink, minv = k, v
		}
	}
	// s[mink] = minv
}

func main() {

}
