package bloomfilter

import "github.com/daysleep666/someproject/datastruct/bitmap/bitmap"

type BloomFilter struct {
	bm *bitmap.BitMap
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{}
}

func (bl *BloomFilter) Set(n int) {
	bl.bm.Set(bl.hash1(n))
	bl.bm.Set(bl.hash2(n))
	bl.bm.Set(bl.hash3(n))
}

func (bl *BloomFilter) IsExist(n int) bool {
	return bl.bm.IsExist(bl.hash1(n)) && bl.bm.IsExist(bl.hash2(n)) && bl.bm.IsExist(bl.hash3(n))
}

func (bl *BloomFilter) hash1(n int) int {
	return n % 8
}

func (bl *BloomFilter) hash2(n int) int {
	return (n * n) % 8
}

func (bl *BloomFilter) hash3(n int) int {
	return (n + n) % 8
}
