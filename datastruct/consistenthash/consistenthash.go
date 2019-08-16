package consistenthash

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"github.com/spaolacci/murmur3"
)

type cirNode struct {
	index uint32
	value string
}

type cirNodes []*cirNode

func (u cirNodes) Less(i, j int) bool { return u[i].index > u[j].index }

func (u cirNodes) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u cirNodes) Len() int { return len(u) }

const defaultVirNode = 200

// Hash ConsistentHash 结构
type Hash struct {
	nodes        map[uint32]string // 辅助节点，不参与具体业务
	cirNodes     *atomic.Value     // []{index在环上的位置,value指向的值}
	virNodeCount int               //  虚拟节点数量
	mu           *sync.Mutex
}

// New 返回hash的指针
func New(virNodeCount int) *Hash {
	if virNodeCount <= 0 {
		virNodeCount = defaultVirNode
	}

	return &Hash{
		nodes:        make(map[uint32]string),
		virNodeCount: virNodeCount,
		cirNodes: func() *atomic.Value {
			cn := make(cirNodes, 0, defaultVirNode*2)
			v := new(atomic.Value)
			v.Store(cn)
			return v
		}(),
		mu: new(sync.Mutex),
	}
}

// Get 通过key 获取对应缓存地址
// 通过key获取数字hash值
// 根据hash值 查找circle内 顺时针右边的node节点
// 根据node节点 获取对应的addr
func (c *Hash) Get(key string) (string, error) {
	cn := c.load()
	if len(cn) == 0 {
		return "", errors.WithStack(fmt.Errorf("empty pool"))
	}
	if len(key) == 0 {
		return "", errors.WithStack(fmt.Errorf("key is nil"))
	}
	hashK := c.hashKey(key)

	v := sort.Search(len(cn), func(x int) bool {
		return cn[x].index > hashK
	})

	// If there is no such index, Search returns n.
	if v == len(cn) {
		v = 0
	}
	return cn[v].value, nil
}

func (c *Hash) hashKey(key string) uint32 { // murmur3
	return murmur3.Sum32([]byte(key))
}

// Add 初始化时设置
// 此处加锁考虑多线程操作对数据操作
// set 流程
// 1 将现有地址 hash后取值 获得circle内对应的数字点
// 2 将hash后的值与addr做成映射NewRedisKvClient
// 3 nodes 保存现有存在的点 将nNewRedisKvClients排序方便get时查找
func (c *Hash) Add(value string) {
	if c.isExist(value) {
		return
	}
	if value == "" {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < c.virNodeCount; i++ {
		virKey := c.hashKey(generateKey(value, i))
		c.nodes[virKey] = value
	}
	c.updateNodes()
}

// nodes是根据circle 的key生成的
func (c *Hash) updateNodes() {
	var tmpCN cirNodes
	for k, v := range c.nodes {
		tmpCN = append(tmpCN, &cirNode{index: k, value: v})
	}
	sort.Sort(tmpCN)
	c.store(tmpCN)
}

// 是否存在这个值
func (c *Hash) isExist(value string) bool {
	virKey := c.hashKey(generateKey(value, 0))
	_, isExist := c.nodes[virKey]
	return isExist
}

// Del 删除节点
func (c *Hash) Del(value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := 0; i < c.virNodeCount; i++ {
		virKey := c.hashKey(generateKey(value, i))
		delete(c.nodes, virKey)
	}
	c.updateNodes()
}

func (c *Hash) load() cirNodes {
	return c.cirNodes.Load().(cirNodes)
}

func (c *Hash) store(cn cirNodes) {
	c.cirNodes.Store(cn)
}

func generateKey(value string, index int) string {
	return fmt.Sprintf("%s#%d", value, index)
}
