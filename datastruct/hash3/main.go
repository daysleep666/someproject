package main

import (
	"fmt"
)

// 哈希函数： hash(key) = hashA + hashB
// 解决冲突：拉链法
// 扩容：旧桶慢慢转移到新桶

// 1.将hashmap的整体思路是实现下
// 2.试验下指针偏移选址

const (
	bucketCnt     = 1 << 3
	loadFactorNum = 13
	loadFactorDen = 2
)

type kv struct { // 键值对
	key   string
	value string
}

type topHash struct {
	topHash int // 高8位
	kv      *kv // 键值对指针
}

type bucket struct {
	overFlowBucket *bucket             // 溢出桶的指针
	tophash        [bucketCnt]*topHash // 键值对组
}

type myMap struct {
	buckets        []*bucket // 新桶
	oldBuckets     []*bucket // 旧桶
	curCount       int       // 当前数量
	curBucketCount int       // 当前桶(包括溢出桶)的数量
	B              uint
}

func newBucket() *bucket {
	b := &bucket{
		tophash:        [bucketCnt]*topHash{},
		overFlowBucket: nil,
	}
	for i, _ := range b.tophash {
		b.tophash[i] = &topHash{}
	}
	return b
}

func NewMap(_maxSize int) *myMap {
	var B uint = 1
	for overLoadFactor(_maxSize, B) {
		B++
	}
	m := &myMap{
		buckets:    make([]*bucket, 1<<B),
		oldBuckets: nil,
		B:          B,
	}
	for i, _ := range m.buckets {
		m.buckets[i] = newBucket()
	}

	for _, v := range m.buckets {
		fmt.Println(v)
	}
	return m
}

func (m *myMap) hash(_key string) (high int, low int) {
	// 会将key的哈希分为高8位和低8位
	// 低8位用来寻找是哪个桶
	// 高8位用来在桶里寻找是哪个值

	// 这里我用time33的hash函数
	// 主要是表示个大概意思，golang里用的hash函数可以参考runtime.alg.go

	var hash int
	for _, v := range _key {
		hash += int(v) * 33
	}
	high = hash & 0x7FFFFFFF                // 代表高8位
	low = (hash & 0x7FFFFFFF) % int(1<<m.B) // 代表低8位
	return
}

// 是否正在扩容
func (m *myMap) isGrowing() bool {
	return m.oldBuckets != nil
}

// 是否超过预定的装载因子
func overLoadFactor(_count int, _b uint) bool {
	// return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
	// 2^B * 6.5
	return _count > bucketCnt && _count > loadFactorNum*(2<<_b/loadFactorDen)
}

// 是否超过预定的溢出桶的数量
func (m *myMap) tooManyOverFloawBuckets() bool {
	// return noverflow >= uint16(1)<<(B&15)
	return m.curBucketCount >= 1<<m.B&0111
}

// 是否已经被迁移了
func (m *myMap) evacuated(_bucket *bucket) bool {
	if _bucket == nil {
		return true
	}
	return _bucket.tophash[0].topHash == -1
}

// 迁移一个桶
func (m *myMap) evacuate(_bucket *bucket) {
	// tmpBucket := _bucket
	// for tmpBucket != nil {
	// 	for _, v := range tmpBucket.tophash {
	// 		key := v.kv.key

	// 	}
	// 	tmpBucket = tmpBucket.overFlowBucket
	// }
}

// 扩容
func (m *myMap) hashGrow() {
	var bigger uint = 1
	// 1<<B
	oldBucket := m.oldBuckets
	newBucket := make([]*bucket, 1<<(m.B+bigger))

	m.curCount = 0
	m.curBucketCount = 0
	m.buckets = newBucket
	m.oldBuckets = oldBucket
}

func (m *myMap) GetValue(_key string) (string, bool) {
	high, low := m.hash(_key)
	bucket := m.buckets[low]
	// 先看看旧桶是否被迁移到了新桶里
	oldBucket := m.oldBuckets[low]
	if !m.evacuated(oldBucket) { // 返回false没有被迁移，那在旧桶里找
		bucket = oldBucket
	}
	for bucket != nil {
		for _, v := range bucket.tophash {
			// 用高8位来在找到键值对的位置
			// 再golang源码里，这个其实是一个及其优化的过程
			// 键值对的存储方式是 key1key2....key8value1value2...value8
			// 寻找key和value是通过指针的偏移得到的，这个我会在后面单独写下。
			if v.topHash == high && v.kv.key == _key {
				return v.kv.value, true
			}
		}
		bucket = bucket.overFlowBucket // 在当前桶里没找到，进入溢出桶接着找
	}

	return "", false
}

func (m *myMap) SetValue(_key string, _value string) {
	// 先判断下是否正在迁移过程中
	// 如果正在迁移过程中，那就将这个key对应的旧桶迁移到新桶
	high, low := m.hash(_key)
	bucket := m.buckets[low]

	if m.isGrowing() {
		// 将旧桶搬迁到新桶
		m.evacuate(m.oldBuckets[low])
	}

	// 总体来说就是先在桶里找存在不存在这个key，
	// 如果存在就更新值，如果不存在就放入桶的空余位置
	// 或者新增一个桶

	for bucket != nil {
		for _, v := range bucket.tophash {
			if v.kv != nil && v.topHash == high && v.kv.key == _key {
				v.kv.value = _value
				return
			}
		}
		if bucket.overFlowBucket == nil {
			break
		}
		bucket = bucket.overFlowBucket
	}
	var isFull bool = true
	// 在桶中没有这个值，需要新增了
	for i := 0; i < bucketCnt; i++ {
		if bucket.tophash[i] == nil {
			bucket.tophash[i] = &topHash{topHash: high, kv: &kv{key: _key, value: _value}}
			isFull = false
			break
		}
	}

	// 桶里没地方，新增一个桶
	if isFull {
		bucket.overFlowBucket = newBucket()
		bucket.overFlowBucket.tophash[0] = &topHash{topHash: high, kv: &kv{key: _key, value: _value}}
		m.curBucketCount++
	}
	m.curCount++

	// 没有在扩容中，过大的装载因子或者过多的溢出桶都会导致迁移发生
	if !m.isGrowing() && (overLoadFactor(m.curCount, m.B) || m.tooManyOverFloawBuckets()) {
		m.hashGrow()
	}

}

func main() {
	var m = NewMap(10)
	for i := 0; i < 2; i++ {
		key := fmt.Sprintf("%v", i)
		value := fmt.Sprintf("%v", i)
		m.SetValue(key, value)
	}
}
