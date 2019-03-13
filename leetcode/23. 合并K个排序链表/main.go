package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	nums := make([]int, 0)
	tmp := &ListNode{}
	for _, v := range lists {
		tmp = v
		for tmp != nil {
			nums = append(nums, tmp.Val)
			tmp = tmp.Next
		}
	}
	if len(nums) == 0 {
		return nil
	}

	var head *ListNode
	heapify(nums, 0, len(nums)-1)
	head = &ListNode{
		Val: nums[0],
	}

	tmpNode := head
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i-1)
		tmpNode.Next = &ListNode{
			Val: nums[0],
		}
		tmpNode = tmpNode.Next
	}
	return head
}

func heapify(nums []int, low, high int) {
	for i := (high - 1) / 2; i >= low; i-- {
		down(nums, i, high)
	}
}

func down(nums []int, low, high int) {
	i := low
	for {
		childIndex := i
		if i*2+1 <= high && nums[i] < nums[i*2+1] {
			childIndex = i*2 + 1
		}
		if i*2+2 <= high && nums[childIndex] < nums[i*2+2] {
			childIndex = i*2 + 2
		}
		if childIndex == i {
			break
		}
		nums[i], nums[childIndex] = nums[childIndex], nums[i]
	}
}

//----------

func mergeKLists1(lists []*ListNode) *ListNode {
	for i := 1; i < len(lists); i++ {
		lists[0] = mergeTwoList(lists[0], lists[i])
	}
	return lists[0]
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	l3 := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			fmt.Println("l1:", l1.Val)
			l3.Next = l1
			l1 = l1.Next
		} else {
			fmt.Println("l2:", l2.Val)
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 == nil {
		l3.Next = l2
	} else {
		l3.Next = l1
	}
	dummyHead = dummyHead.Next
	return dummyHead
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: -1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	mergeTwoList(l1, l2)
	for l1 != nil {
		fmt.Printf("%v ", l1.Val)
		l1 = l1.Next
	}
}
