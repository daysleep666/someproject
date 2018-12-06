package main

import (
	"fmt"
	"math/rand"
	"time"
)

// #define ZSKIPLIST_MAXLEVEL 32
// #define ZSKIPLIST_P 0.25

// typedef struct zskiplistNode {
//     robj *obj;
//     double score;
//     struct zskiplistNode *backward;
//     struct zskiplistLevel {
//         struct zskiplistNode *forward;
//         unsigned int span;
//     } level[];
// } zskiplistNode;

// typedef struct zskiplist {
//     struct zskiplistNode *header, *tail;
//     unsigned long length;
//     int level;
// } zskiplist;

const (
	ZSKIPLIST_MAXLEVEL = 32
	ZSKIPLIST_P        = 25
)

type ZskipListLevel struct {
	Forward *ZskipListNode
	Span    uint // 在每层到下一个节点跨越的长度，相对长度
}

type ZskipListNode struct {
	Member   string         // 在redis中是obj的数据类型，但是这里主要是为了说明redis中的跳表，所以简化为string
	Score    int            // score
	BackWard *ZskipListNode // 后端指针
	Level    []*ZskipListLevel
}

type ZskipList struct {
	Header *ZskipListNode // 头节点
	Tail   *ZskipListNode // 尾节点
	Length uint           // 节点数量
	Level  int            // 最高层数
}

func GetNewZskipList() *ZskipList {
	return &ZskipList{
		Header: getNewNode(ZSKIPLIST_MAXLEVEL),
		Level:  1,
	}
}

func getNewNode(_level int) *ZskipListNode {
	var znode = &ZskipListNode{Level: make([]*ZskipListLevel, _level)}
	for i, _ := range znode.Level {
		znode.Level[i] = &ZskipListLevel{}
	}
	return znode
}

func (zln *ZskipList) Insert(_member string, _newScore int) {
	var (
		updateNode = make([]*ZskipListNode, ZSKIPLIST_MAXLEVEL)
		tmpNode    *ZskipListNode
		frontNode         = zln.Header
		rank       []uint = make([]uint, ZSKIPLIST_MAXLEVEL+1)
	)
	for i := zln.Level - 1; i >= 0; i-- {
		rank[i] += rank[i+1]
		tmpNode = frontNode.Level[i].Forward
		for tmpNode != nil && tmpNode.Score < _newScore {
			frontNode = tmpNode
			rank[i] += tmpNode.Level[i].Span
			tmpNode = tmpNode.Level[i].Forward
		}
		updateNode[i] = frontNode
	}

	// 产生了新层
	var randomLevel = zln.getRandomLevel()
	if zln.Level < randomLevel {
		for i := zln.Level; i < randomLevel; i++ {
			updateNode[i] = zln.Header
			updateNode[i].Level[i].Span = zln.Length
		}
		zln.Level = randomLevel
	}

	var newNode = getNewNode(randomLevel)
	newNode.Member = _member
	newNode.Score = _newScore
	// 塞个后退指针
	newNode.BackWard = updateNode[0]
	if updateNode[0].Level[0].Forward != nil {
		updateNode[0].Level[0].Forward.BackWard = newNode
	} else {
		zln.Tail = newNode
	}

	for i := 0; i < randomLevel; i++ {
		newNode.Level[i].Forward = updateNode[i].Level[i].Forward
		newNode.Level[i].Span = updateNode[i].Level[i].Span - (rank[0] - rank[i])
		updateNode[i].Level[i].Span = rank[0] + 1 - rank[i]
		updateNode[i].Level[i].Forward = newNode
	}
	for i := randomLevel; i < zln.Level; i++ {
		updateNode[i].Level[i].Span++
	}

	zln.Length++ // 长度+1
}

func (zln *ZskipList) Output() {
	var (
		tmpNode *ZskipListNode
	)

	for i := zln.Level - 1; i >= 0; i-- {
		fmt.Printf("Level %v   :", i)
		tmpNode = zln.Header
		for tmpNode != nil {
			fmt.Printf("  [member: %v, span: %v]", tmpNode.Member, tmpNode.Level[i].Span)
			// fmt.Printf("  member: %v, score: %v, span: %v", tmpNode.Member, tmpNode.Score, tmpNode.Level[i].Span)
			tmpNode = tmpNode.Level[i].Forward
		}
		fmt.Println()
	}
	fmt.Printf("NodeLength: %v\n", zln.Length)
}

func (zln *ZskipList) Display() {
	var tmpNode = zln.Header.Level[0].Forward
	for tmpNode != nil {
		fmt.Printf("member = %v, score = %v, rank = %v \n", tmpNode.Member, tmpNode.Score, tmpNode.Level[0].Span)
		tmpNode = tmpNode.Level[0].Forward
	}
}

func (zln *ZskipList) DisplayBackWard() {
	fmt.Printf("DisplayBackWard:")
	var tmpNode = zln.Tail
	for tmpNode != nil {
		fmt.Printf("%v ", tmpNode.Score)
		tmpNode = tmpNode.BackWard
	}
	fmt.Println()
}

func (zln *ZskipList) getRandomLevel() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < ZSKIPLIST_MAXLEVEL; i++ {
		randValue := r.Intn(100)
		if randValue > ZSKIPLIST_P {
			return i + 1
		}
	}
	return ZSKIPLIST_MAXLEVEL
}

func main() {
	// s := rand.NewSource(time.Now().Unix())
	// r := rand.New(s)
	var zskipListNode = GetNewZskipList()
	for i := 0; i < 6; i++ {
		zskipListNode.Insert(fmt.Sprintf("%v", i), rand.Intn(100))
	}
	// zskipListNode.Display()
	zskipListNode.DisplayBackWard()
	zskipListNode.Output()
}
