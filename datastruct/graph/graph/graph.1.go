package graph

import "fmt"

type MGraph struct {
	Nodes []int
	M     map[int][]int
}

func NewGraph() *MGraph {
	return &MGraph{Nodes: make([]int, 0), M: make(map[int][]int)}

}

func (m *MGraph) Find(n int) int {
	for i, v := range m.Nodes {
		if v == n {
			return i
		}
	}
	return -1
}

func (m *MGraph) Push(n1, n2 int) {
	if m.Find(n1) == -1 {
		m.Nodes = append(m.Nodes, n1)
	}
	if m.Find(n2) == -1 {
		m.Nodes = append(m.Nodes, n2)
	}
	if _, isExist := m.M[n1]; !isExist {
		m.M[n1] = make([]int, 0, 1)
	}
	for _, v := range m.M[n1] {
		if v == n2 {
			return
		}
	}
	m.M[n1] = append(m.M[n1], n2)
	for _, v := range m.M[n2] {
		if v == n1 {
			return
		}
	}
	m.M[n2] = append(m.M[n2], n1)
}

func (m *MGraph) GetChildred(v int) []int {
	return m.M[v]
}

func (m *MGraph) BFS(from int) {
	visited := make(map[int]bool)
	list := make([]int, 0)
	list = append(list, from)
	visited[from] = true
	fmt.Printf("%v", from)
	for {
		if len(list) == 0 {
			break
		}
		tmpList := make([]int, 0)
		for _, v := range list {
			children := m.GetChildred(v)
			for _, c := range children {
				if !visited[c] {
					tmpList = append(tmpList, m.GetChildred(c)...)
					fmt.Printf("->%v", c)
					visited[c] = true
				}
			}
		}
		list = tmpList
	}
	fmt.Println()
}
