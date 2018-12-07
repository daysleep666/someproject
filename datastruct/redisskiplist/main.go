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
	ZSKIPLIST_MAXLEVEL = 32 // 最多32层
	ZSKIPLIST_P        = 25 // 25%的几率提升
)

type MemberStruct struct {
	MemberName    string
	MemberContent *MemberContentStruct
}

type MemberContentStruct struct {
	Score      int64
	UpdateTime int64
}

//  mymember < othermember  = -1
//  mymember == othermember = 0
//  mymember > othermember  = 1

func (myMember *MemberStruct) Cmp(_otherMemeber *MemberStruct) int {
	if myMember.MemberName == _otherMemeber.MemberName {
		return 0
	}

	// 优先比较分数 越大越好
	if myMember.MemberContent.Score > _otherMemeber.MemberContent.Score {
		return 1
	} else if myMember.MemberContent.Score < _otherMemeber.MemberContent.Score {
		return -1
	}

	// 然后比较更新时间 越小越好
	if myMember.MemberContent.UpdateTime < _otherMemeber.MemberContent.UpdateTime {
		return 1
	} else if myMember.MemberContent.UpdateTime > _otherMemeber.MemberContent.UpdateTime {
		return -1
	}

	// 最后毕竟名字 ascii越小 越靠前
	if myMember.MemberName > _otherMemeber.MemberName {
		return 1
	} else if myMember.MemberName < _otherMemeber.MemberName {
		return -1
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
	BackWard *ZskipListNode // 后端指针
	Level    []*ZskipListLevel
}

func (node *ZskipListNode) GetScore() int64 {
	return node.Member.MemberContent.Score
}

type ZskipList struct {
	Header *ZskipListNode                  // 头节点
	Tail   *ZskipListNode                  // 尾节点
	Length int                             // 节点数量
	Level  int                             // 最高层数
	Dict   map[string]*MemberContentStruct // member->score 映射
}

func GetNewZskipList() *ZskipList {
	return &ZskipList{
		Header: getNewNode(ZSKIPLIST_MAXLEVEL, "", -1, -1),
		Level:  1,
		Dict:   make(map[string]*MemberContentStruct),
	}
}

func getNewNode(_level int, _memberName string, _score int64, _updateTime int64) *ZskipListNode {
	var znode = &ZskipListNode{Level: make([]*ZskipListLevel, _level), Member: &MemberStruct{MemberName: _memberName, MemberContent: &MemberContentStruct{Score: _score, UpdateTime: _updateTime}}}
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

func (zln *ZskipList) GetMemberStructByName(_memberName string) *MemberStruct {
	memberContent := zln.Dict[_memberName]
	if memberContent == nil {
		return nil
	}
	return &MemberStruct{MemberName: _memberName, MemberContent: memberContent}
}

func (zln *ZskipList) Insert(_member *MemberStruct) error {
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
			} else if tmpNode.Member.Cmp(_member) == -1 {
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

	var newNode = getNewNode(randomLevel, _member.MemberName, _member.MemberContent.Score, _member.MemberContent.UpdateTime)
	zln.Dict[_member.MemberName] = _member.MemberContent
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
	memberStruct := zln.GetMemberStructByName(_memberName)
	if memberStruct == nil {
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
			if tmpNode.Member.Cmp(memberStruct) == -1 {
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
	zln.Dict[_memberName] = nil
	zln.Length--
	return
}

func (zln *ZskipList) FindRank(_memberName string) int {
	memberStruct := zln.GetMemberStructByName(_memberName)
	if memberStruct == nil {
		return 0
	}
	var (
		tmpNode   *ZskipListNode
		frontNode       = zln.Header
		rank      []int = make([]int, ZSKIPLIST_MAXLEVEL+1)
	)

	for i := zln.Level - 1; i >= 0; i-- {
		rank[i] = rank[i+1]
		tmpNode = frontNode.Level[i].Forward
		for tmpNode != nil {
			if tmpNode.Member.Cmp(memberStruct) == -1 {
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

func (zln *ZskipList) FindRankFromTo(_from, _to int) []*MemberStruct { //[from,to)
	if _from >= _to || _from >= zln.Length || _from <= 0 {
		return nil
	}

	var end = _to
	if _to > zln.Length {
		end = zln.Length
	}
	var (
		totalSpan int
		tmpNode   = zln.Header
		frontNode = zln.Header
		result    = make([]*MemberStruct, end-_from+1)
	)

	for i := zln.Level - 1; i >= 0; i-- {
		tmpNode = frontNode.Level[i].Forward
		for tmpNode != nil {
			var tmpSpan = totalSpan + frontNode.Level[i].Span // 当前节点的rank
			if tmpSpan < _from {                              // 继续累加
				totalSpan = tmpSpan
				tmpNode = tmpNode.Level[i].Forward
				continue
			} else if tmpSpan > end { // 超出去了 直接跳出去
				break
			}
			// 这个是想找的那个
			// 先累加前面的

			var tmpFrontNode = tmpNode
			for i := tmpSpan; i >= _from; i-- {
				result[i-_from] = tmpFrontNode.Member
				tmpFrontNode = tmpFrontNode.BackWard
			}
			// 再累加后面的
			for i := tmpSpan + 1; i <= end; i++ {
				result[i-_from] = tmpNode.Member
				tmpNode = tmpNode.Level[0].Forward
			}
			return result
		}
	}
	return nil
}

func (zln *ZskipList) Output() {
	var (
		tmpNode *ZskipListNode
	)

	for i := zln.Level - 1; i >= 0; i-- {
		fmt.Printf("Level %v   :", i)
		tmpNode = zln.Header
		for tmpNode != nil {
			fmt.Printf("  [member: %v, score: %v, span: %v]", tmpNode.Member.MemberName, tmpNode.GetScore(), tmpNode.Level[i].Span)
			tmpNode = tmpNode.Level[i].Forward
		}
		fmt.Println()
	}
	fmt.Printf("NodeLength: %v, NodeLevel: %v\n", zln.Length, zln.Level)
}

func (zln *ZskipList) DisplayBackWard() {
	fmt.Printf("DisplayBackWard:")
	var tmpNode = zln.Tail
	for tmpNode != nil {
		fmt.Printf("%v ", tmpNode.GetScore())
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
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	var zskipListNode = GetNewZskipList()
	for i := 0; i < 10; i++ {
		zskipListNode.Insert(&MemberStruct{MemberName: fmt.Sprintf("user:%v", i), MemberContent: &MemberContentStruct{UpdateTime: time.Now().Unix(), Score: r.Int63n(10000)}})
	}

	zskipListNode.Output()

	fmt.Println("--------------------------")
	// zskipListNode.Delete("1")
	// zskipListNode.Delete("2")
	// zskipListNode.Delete("0")
	// zskipListNode.Insert(&MemberStruct{MemberName: "hey", UpdateTime: time.Now().Unix()}, 88)
	// zskipListNode.Display()
	zskipListNode.DisplayBackWard()
	// zskipListNode.Output()
	// zskipListNode.DisplayRank()
	st := time.Now().UnixNano()
	rank := zskipListNode.FindRank("user:666")
	fmt.Printf("\nrank is %v, spend %v ms\n", rank, (time.Now().UnixNano()-st)/1000000)

	result := zskipListNode.FindRankFromTo(1, 10)
	for i, v := range result {
		fmt.Printf("rank:%v member:%v score:%v \n", i+1, v.MemberName, v.MemberContent.Score)
	}
}
