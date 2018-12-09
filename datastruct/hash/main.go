package main

import (
	"fmt"
)

// 除留取余-》拉链法-》哈希桶

// 函数函数使用除留余数法
// f( key ) = key mod p ( p ≤ m )
// 哈希冲突使用线性探测在散列

type keyValue struct {
	Key   string
	Value interface{}
}

type myMap struct {
	// hashTable = make([]*keyValue, M)
	HashTable     []*keyValue
	LoadingFactor float32 // 装填因子
	MaxSize       uint    // 最大长度
	PrimeNumber   uint    // 接近最大长度的质数
}

func NewMap(_maxSize uint) *myMap {
	if _maxSize == 0 {
		panic("maxsize must more than zero")
	}
	return &myMap{
		HashTable:     make([]*keyValue, _maxSize),
		LoadingFactor: 0.0,
		MaxSize:       _maxSize,
		PrimeNumber:   (findPrimeNumber(_maxSize)),
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

func (m *myMap) GetValue(_key string) (interface{}, bool) {
	hash, isExist := m.reHash(_key)
	if isExist {
		return m.HashTable[hash].Value, true
	}
	return 0, false
}

func (m *myMap) SetVaue(_key string, _value interface{}) {
	hash, _ := m.reHash(_key)
	if hash == -1 {
		return
	}
	m.HashTable[hash] = &keyValue{Key: _key, Value: _value}
}

func main() {
	m := NewMap(10)
	for i := 0; i < 11; i++ {
		m.SetVaue(fmt.Sprintf("%v", i), i)
	}
	for i := 0; i < 11; i++ {
		fmt.Println(m.GetValue(fmt.Sprintf("%v", i)))
	}
	m.SetVaue("key", "hello hashtable")
	fmt.Println(m.GetValue("key"))
}
