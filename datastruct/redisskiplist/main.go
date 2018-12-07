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

type MemberStruct struct {
	MemberName string
	UpdateTime int64
}

//  mymember < othermember  = -1
//  mymember == othermember = 0
//  mymember > othermember  = 1

func (myMember *MemberStruct) Cmp(_otherMemeber *MemberStruct) int {
	if myMember.MemberName == _otherMemeber.MemberName {
		return 0
	}
	if myMember.UpdateTime < _otherMemeber.UpdateTime {
		return -1
	} else if myMember.UpdateTime > _otherMemeber.UpdateTime {
		return 1
	} else {
		if myMember.MemberName < _otherMemeber.MemberName {
			return -1
		} else if myMember.MemberName > _otherMemeber.MemberName {
			return 1
		}
	}
	return 0
}

func (myMember *MemberStruct) IsSame(_otherMemberName string) bool {
	return myMember.MemberName == _otherMemberName
}

type ZskipListLevel struct {
	Forward *ZskipListNode
	Span    int // 在每层到下一个节点跨越的长度，相对长度
}

type ZskipListNode struct {
	Member   *MemberStruct  // 在redis中是obj的数据类型，但是这里主要是为了说明redis中的跳表，所以简化为string
	Score    int            // score
	BackWard *ZskipListNode // 后端指针
	Level    []*ZskipListLevel
}

type ZskipList struct {
	Header *ZskipListNode // 头节点
	Tail   *ZskipListNode // 尾节点
	Length int            // 节点数量
	Level  int            // 最高层数
	Dict   map[string]int // member->score 映射
}

func GetNewZskipList() *ZskipList {
	return &ZskipList{
		Header: getNewNode(ZSKIPLIST_MAXLEVEL),
		Level:  1,
		Dict:   make(map[string]int),
	}
}

func getNewNode(_level int) *ZskipListNode {
	var znode = &ZskipListNode{Level: make([]*ZskipListLevel, _level), Member: &MemberStruct{}}
	for i, _ := range znode.Level {
		znode.Level[i] = &ZskipListLevel{}
	}
	return znode
}

// 总结下span的计算方法
// 1.计算出第i层头节点到updateNode[i]节点的长度，存在rank[i]中
// 2.计算高度
// 2.1 如果计算出的高度大于当前高度，就将大于高度的部分设置为Length
// 2.2 如果计算出的高度小于当前高度，就将大于等于计算出高度到当前高度到部分补一个一。因为下面到层多了一个节点
// 3 rank[0]++ 等于 头节点到新节点的长度
// 4.新节点到下一个节点的span = 新节点的上一个节点的span - (头节点到新节点的长度 - 头节点到新节点上一个节点的长度)
// 5.新节点上一个节点到新节点的span = 头节点到新节点的长度 - 头节点到新节点上一个节点的长度

func (zln *ZskipList) Insert(_member *MemberStruct, _newScore int) error {
	var (
		updateNode = make([]*ZskipListNode, ZSKIPLIST_MAXLEVEL)
		tmpNode    *ZskipListNode
		frontNode        = zln.Header
		rank       []int = make([]int, ZSKIPLIST_MAXLEVEL+1)
	)
	for i := zln.Level - 1; i >= 0; i-- {
		rank[i] = rank[i+1]
		tmpNode = frontNode.Level[i].Forward
		for tmpNode != nil {
			if tmpNode.Member.Cmp(_member) == 0 {
				return fmt.Errorf("%v is exist", _member.MemberName)
			} else if tmpNode.Score < _newScore || (tmpNode.Score == _newScore && tmpNode.Member.Cmp(_member) == -1) {
				rank[i] += frontNode.Level[i].Span //这里注意，记录的是跨越了多少节点到达这里
				frontNode = tmpNode
				tmpNode = tmpNode.Level[i].Forward
			} else {
				break
			}
		}
		updateNode[i] = frontNode
	}

	// 产生了新层
	var randomLevel = zln.getRandomLevel()
	if zln.Level < randomLevel {
		for i := zln.Level; i < randomLevel; i++ {
			rank[i] = 0
			updateNode[i] = zln.Header
			updateNode[i].Level[i].Span = zln.Length
		}
		zln.Level = randomLevel
	}

	var newNode = getNewNode(randomLevel)
	newNode.Member = _member
	newNode.Score = _newScore
	zln.Dict[_member.MemberName] = _newScore
	// 塞个后退指针
	newNode.BackWard = updateNode[0]
	if updateNode[0].Level[0].Forward != nil {
		updateNode[0].Level[0].Forward.BackWard = newNode
	} else {
		zln.Tail = newNode
	}

	for i := 0; i < randomLevel; i++ {
		newNode.Level[i].Forward = updateNode[i].Level[i].Forward
		updateNode[i].Level[i].Forward = newNode
		newNode.Level[i].Span = updateNode[i].Level[i].Span - (rank[0] - rank[i])
		updateNode[i].Level[i].Span = rank[0] - rank[i] + 1
	}
	for i := randomLevel; i < zln.Level; i++ {
		updateNode[i].Level[i].Span++
	}

	zln.Length++ // 长度+1
	return nil
}

func (zln *ZskipList) Delete(_memberName string) {
	score, isExist := zln.Dict[_memberName]
	if !isExist { // 不存在 不找了
		return
	}
	var (
		tmpNode    *ZskipListNode
		frontNode  = zln.Header
		updateNode = make([]*ZskipListNode, ZSKIPLIST_MAXLEVEL)
	)

	for i := zln.Level - 1; i >= 0; i-- {
		tmpNode = frontNode.Level[i].Forward
		for tmpNode != nil {
			if tmpNode.Score < score {
				frontNode = tmpNode
				tmpNode = tmpNode.Level[i].Forward
			} else {
				break
			}
		}
		updateNode[i] = frontNode
	}
	// 降级
	var length = zln.Level - 1
	for i := length; i >= 0; i-- {
		if updateNode[i].Level[i].Forward == nil || !updateNode[i].Level[i].Forward.Member.IsSame(_memberName) { // 说明这一层没有这个节点
			updateNode[i].Level[i].Span--
			continue
		}
		updateNode[i].Level[i].Span += updateNode[i].Level[i].Forward.Level[i].Span // 加上要删掉的那个节点到后面节点的跨度
		if updateNode[i].Level[i].Forward.Level[i].Span != 0 {
			updateNode[i].Level[i].Span-- // 后面还有节点，所以需要减去1
		}

		if updateNode[i].Level[i].Span == zln.Length { // 这层没有其他节点了
			zln.Level--
		}
		updateNode[i].Level[i].Forward = updateNode[i].Level[i].Forward.Level[i].Forward
	}
	zln.Length--
	return
}

func (zln *ZskipList) Output() {
	var (
		tmpNode *ZskipListNode
	)

	for i := zln.Level - 1; i >= 0; i-- {
		fmt.Printf("Level %v   :", i)
		tmpNode = zln.Header
		for tmpNode != nil {
			fmt.Printf("  [member: %v, score: %v, span: %v]", tmpNode.Member.MemberName, tmpNode.Score, tmpNode.Level[i].Span)
			// fmt.Printf("  member: %v, score: %v, span: %v", tmpNode.MemberName, tmpNode.Score, tmpNode.Level[i].Span)
			tmpNode = tmpNode.Level[i].Forward
		}
		fmt.Println()
	}
	fmt.Printf("NodeLength: %v, NodeLevel: %v\n", zln.Length, zln.Level)
}

func (zln *ZskipList) Display() {
	var tmpNode = zln.Header.Level[0].Forward
	for tmpNode != nil {
		fmt.Printf("member = %v, score = %v, rank = %v \n", tmpNode.Member, tmpNode.Score, tmpNode.Level[0].Span)
		tmpNode = tmpNode.Level[0].Forward
	}
}

func (zln *ZskipList) DisplayRank() {
	var tmpNode = zln.Header.Level[0].Forward
	var rank int
	for tmpNode != nil {
		rank++
		fmt.Printf("[member = %v, rank = %v] ", tmpNode.Member.MemberName, rank)
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

func (zln *ZskipList) FindRank(_memberName string) int {
	score, isExist := zln.Dict[_memberName]
	if !isExist { // 不存在 不找了
		return 0
	}
	var (
		tmpNode   *ZskipListNode
		frontNode       = zln.Header
		rank      []int = make([]int, ZSKIPLIST_MAXLEVEL+1)
	)

	for i := int(zln.Length - 1); i >= 0; i-- {
		rank[i] = rank[i+1]
		tmpNode = frontNode.Level[i].Forward
		for tmpNode != nil {
			if tmpNode.Score < score {
				rank[i] += frontNode.Level[i].Span
				frontNode = tmpNode
				tmpNode = tmpNode.Level[i].Forward
			} else if tmpNode.Member.IsSame(_memberName) { // 找到了
				return rank[i] + 1
			} else {
				break
			}
		}
	}
	return 0
}

func main() {
	// s := rand.NewSource(time.Now().Unix())
	// r := rand.New(s)
	var zskipListNode = GetNewZskipList()
	// for i := 0; i < 4; i++ {
	// 	zskipListNode.Insert(&MemberStruct{MemberName: fmt.Sprintf("%v", i), UpdateTime: time.Now().Unix()}, rand.Intn(100))
	// }
	zskipListNode.Insert(&MemberStruct{MemberName: "2", UpdateTime: time.Now().Unix()}, 1)
	zskipListNode.Insert(&MemberStruct{MemberName: "1", UpdateTime: time.Now().Unix()}, 1)
	zskipListNode.Insert(&MemberStruct{MemberName: "3", UpdateTime: time.Now().Unix()}, 1)

	zskipListNode.Output()

	fmt.Println("--------------------------")
	// zskipListNode.Delete("1")
	// zskipListNode.Delete("2")
	// zskipListNode.Delete("0")
	// zskipListNode.Insert(&MemberStruct{MemberName: "hey", UpdateTime: time.Now().Unix()}, 88)
	// zskipListNode.Display()
	// zskipListNode.DisplayBackWard()
	zskipListNode.Output()
	// zskipListNode.DisplayRank()
	rank := zskipListNode.FindRank("3")
	fmt.Printf("\nrank is %v\n", rank)
}
