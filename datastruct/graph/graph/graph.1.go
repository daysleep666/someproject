package graph

type MGraph struct {
	Nodes []int
	M     map[int][]int
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
}

func (m *MGraph) BFS() {

}
