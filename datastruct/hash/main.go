package main

import (
	"fmt"
)

// 除留取余-》拉链法-》哈希桶

// 函数函数使用除留余数法
// f( key ) = key mod p ( p ≤ m )
// 哈希冲突使用线性探测在散列

const (
	MAXLOADINGFACTOR = 0.75 // 如果装填因子大于0.75，就需要对哈希表进行扩容
)

type keyValue struct {
	Key   string
	Value interface{}
}

type myMap struct {
	HashTable   []*keyValue
	MaxSize     uint // 最大长度
	CurSize     uint // 当前大小
	PrimeNumber uint // 接近最大长度的质数
}

func NewMap(_maxSize uint) *myMap {
	if _maxSize == 0 {
		panic("maxsize must more than zero")
	}

	// 因为扩容是一件十分耗费性能的事情，
	// 所以再分配的时候需要分配比maxsize多百分之25的空间来避免扩容

	var realMaxSize = uint(float32(_maxSize) / MAXLOADINGFACTOR)

	return &myMap{
		HashTable:   make([]*keyValue, realMaxSize),
		MaxSize:     realMaxSize,
		PrimeNumber: (findPrimeNumber(realMaxSize)),
	}
}

func findPrimeNumber(_num uint) uint { // 找到接近最大长度的质数
	var (
		tmp = _num
	)

	for {
		i := uint(2)
		for ; i < tmp; i++ {
			if tmp%i == 0 {
				break
			}
		}
		if i >= tmp {
			break
		}
		tmp--
	}
	return tmp
}

func (m *myMap) hash(_key string) int { // 除留余数法
	if len(_key) == 0 {
		panic("key length must more than zero")
	}
	return int(_key[0]) % int(m.PrimeNumber)
}

// 线性再探测
func (m *myMap) reHash(_key string) (int, bool) { // 返回的是在数组中的下标，是否存在这个key
	var (
		hash = m.hash(_key)
		tmp  = hash
	)

	if m.HashTable[tmp] != nil && m.HashTable[hash].Key == _key { // 找到了这个值=key
		return hash, true
	}

	for {
		tmp = (tmp + 1) % int(m.MaxSize)
		if m.HashTable[tmp] == nil { // 说明没有这个key
			return tmp, false
		} else if m.HashTable[tmp].Key == _key { // 找到了这个值=key
			return tmp, true
		} else if hash == tmp { // 绕了一圈回来了，说明hashtable满了
			return -1, false
		}
	}
}

func (m *myMap) resize() { // 扩容
	m.MaxSize *= 2                               //容量翻倍
	oldHashTable := m.HashTable                  // 旧的哈希表
	newHashTable := make([]*keyValue, m.MaxSize) //新的哈希表
	m.HashTable = newHashTable
	m.CurSize = 0
	m.PrimeNumber = findPrimeNumber(m.MaxSize) // 新的质数因子
	// 因为要更换质数因子，所以需要处理下之前等数据
	// 扩容是一件十分十分耗费性能的事情
	// 这个操作的时间复杂度是O(n),空间复杂度是 O(n)
	for _, v := range oldHashTable {
		if v == nil {
			continue
		}
		m.SetVaue(v.Key, v.Value)
	}
}

func (m *myMap) Display() {
	fmt.Printf("总大小:%v\n", m.MaxSize)
	fmt.Printf("当前大小:%v\n", m.CurSize)
	fmt.Printf("质数为:%v\n", m.PrimeNumber)
	fmt.Printf("装载因子:%v\n", float32(m.CurSize)/float32(m.MaxSize))
}

func (m *myMap) GetValue(_key string) (interface{}, bool) {
	hash, isExist := m.reHash(_key)
	if isExist {
		return m.HashTable[hash].Value, true
	}
	return 0, false
}

func (m *myMap) SetVaue(_key string, _value interface{}) {
	hash, isExist := m.reHash(_key)
	if hash == -1 {
		fmt.Printf("err:满了")
		return
	}
	m.HashTable[hash] = &keyValue{Key: _key, Value: _value}
	if !isExist {
		m.CurSize++
		curLoadingFactor := float32(m.CurSize) / float32(m.MaxSize)
		if curLoadingFactor > MAXLOADINGFACTOR { // 需要扩容了
			m.resize()
		}
	}
}

func main() {
	m := NewMap(10)
	for i := 0; i < 101; i++ {
		m.SetVaue(fmt.Sprintf("%v", i), i)
	}
	m.Display()
	fmt.Println("-----")
	for i := 0; i < 11; i++ {
		fmt.Println(m.GetValue(fmt.Sprintf("%v", i)))
	}
	m.SetVaue("key", "hello hashtable")
	fmt.Println(m.GetValue("key"))
}
