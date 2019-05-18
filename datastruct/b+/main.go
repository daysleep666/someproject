package main

func main() {

}

//--------------------------

const M = 3

type BPLUSTREEENODE struct {
	Key      int               // 索引值
	Children []*BPLUSTREEENODE // 子树
	// leaf
	Value    int             // 真正的数据
	NextNode *BPLUSTREEENODE // 指向相邻的叶子节点
	IsLeaf   bool            // 是否是叶子节点
}

func NewNode() *BPLUSTREEENODE {
	return &BPLUSTREEENODE{Children: make([]*BPLUSTREEENODE, 0)}
}

func (n *BPLUSTREEENODE) Insert(key int) *BPLUSTREEENODE {
	newNode := NewNode()
	newNode.Key = key
	for i, v := range n.Children {
		if key < v.Key {
			tmp := append(n.Children[:i], newNode)
			tmp = append(tmp, n.Children[i:]...)
			n.Children = tmp
			return newNode
		}
	}
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *BPLUSTREEENODE) IsFull() bool {
	return len(n.Children) > M
}

func Insert(head *BPLUSTREEENODE, value, key int) *BPLUSTREEENODE {
	if head == nil {
		head := NewNode()
		head.Key = key
		node := head.Insert(key)
		node.Value = value
		node.IsLeaf = true
		return head
	}

	parentNodeList := make([]*BPLUSTREEENODE, 0) // 记录每一层的父节点
	node := head
	for !node.IsLeaf {
		parentNodeList = append(parentNodeList, node)
		for _, child := range node.Children {
			if child.IsLeaf { // 应该插在这个节点下面
				break
			}
			if key <= child.Key { // 往下接着找
				node = child
				break
			}
		}
	}

	leafNode := node.Insert(key)
	leafNode.Value = value
	leafNode.IsLeaf = true

	parentIndex := len(parentNodeList) - 1
	for {
		node = parentNodeList[parentIndex]
		if !node.IsFull() {
			return head
		}
		parentIndex--
		if parentIndex == 0 {
			break
		}
		newParentNode := parentNodeList[parentIndex]
		newIndex := (M + 1) / 2
		newChildren := node.Children[:newIndex]
		nnn := newParentNode.Insert(newChildren[newIndex-1].Key)
		nnn.Children = newChildren
		node.Children = node.Children[newIndex:]
	}

	newHead := NewNode()
	newHead.Key = node.Key
	newIndex := (M + 1) / 2
	newChildren := node.Children[:newIndex]
	nnn := NewNode()
	nnn.Key = newChildren[newIndex-1].Key
	nnn.Children = newChildren
	node.Children = node.Children[newIndex:]
	newHead.Children = append(newHead.Children, nnn, node)
	return newHead

	// 还差一个更新当前最大值
}
