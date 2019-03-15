package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	mid := findMid(head)
	if mid == nil || head == nil || head.Next == nil {
		return head
	}
	tmp := mid.Next
	mid.Next = nil
	a := sortList(head)
	b := sortList(tmp)
	return merge(a, b)
}

func merge(a, b *ListNode) *ListNode {
	c := &ListNode{}
	head := c
	for a != nil && b != nil {
		if a.Val < b.Val {
			c.Next = a
			a = a.Next
		} else {
			c.Next = b
			b = b.Next
		}
		c = c.Next
	}
	if a != nil {
		c.Next = a
	}
	if b != nil {
		c.Next = b
	}
	return head.Next
}

func findMid(start *ListNode) *ListNode {
	low, fast := start, start
	for {
		if fast == nil {
			return low
		}
		fast = fast.Next
		if fast == nil {
			return low
		}
		fast = fast.Next
		if fast == nil {
			return low
		}
		low = low.Next
	}
	return low
}

func add(head *ListNode, v int) *ListNode {
	if head == nil {
		head = &ListNode{Val: v}
		return head
	}
	head.Next = add(head.Next, v)
	return head
}

func display(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	display(head.Next)
}

func main() {
	var head *ListNode
	head = add(head, 3)
	head = add(head, 1)
	head = add(head, 4)
	head = add(head, 2)
	head = sortList(head)
	display(head)

}
