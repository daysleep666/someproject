package consistenthash

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 consistenthash.go consistenthash_test.go
// 测试结果是 m个节点,n个节点生成的一致性哈希环.x个key落在这个环上的占比
func TestNewCH(t *testing.T) {
	assert := assert.New(t)

	rand.Seed(time.Now().UnixNano())
	m := make(map[string]int, 0)
	ch := New(300)
	for i := 0; i < 10; i++ {
		ch.Add(fmt.Sprintf("%v", i))
	}
	ch.Del("0")
	sum := 0
	for i := 0; i < 100000; i++ {
		key := generateRandomKey()
		id, err := ch.Get(key)
		assert.Nil(err)
		m[id]++
		sum++
	}

	for k, v := range m {
		fmt.Printf("%v:%.1f%% \n", k, float64(v)/float64(sum)*float64(100))
	}
}

func generateRandomKey() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key := r.Int()
	return fmt.Sprintf("%v", key)
}

/*
murmur3.Sum32的结果
=== RUN   TestNewCH
6:10.8%
3:11.7%
4:11.2%
7:10.9%
8:10.5%
1:11.5%
2:10.9%
9:10.8%
5:11.6%
--- PASS: TestNewCH (1.08s)
PASS

crc32.ChecksumIEEE的结果
=== RUN   TestNewCH
3:14.5%
2:12.7%
7:12.7%
5:9.0%
6:11.9%
9:9.5%
1:8.3%
8:11.6%
4:9.9%
--- PASS: TestNewCH (1.19s)
PASS
ok  	command-line-arguments	1.189s
*/

//------------

//压力测试
func Benchmark_Get(b *testing.B) {
	b.StopTimer()
	h := New(300)
	for i := 0; i < 5; i++ {
		h.Add(fmt.Sprintf("%d", i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		h.Get(fmt.Sprintf("%d", i))
	}
}

// go test -run="Benchmark_Get" -test.bench=".*"
// 改之前的压测
// Benchmark_Get-4   	10000000	       240 ns/op
// 改之后的压测
// Benchmark_Get-4   	10000000	       170 ns/op
