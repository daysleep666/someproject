package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMid(head)
	mid = reverse(mid)

	tmp := head.Next
	tmpHead := head
	for tmp != nil && mid != nil {
		tmpHead.Next = mid
		tmpHead = tmpHead.Next
		mid = mid.Next
		tmpHead.Next = tmp
		if tmp == nil {
			break
		}
		tmpHead = tmpHead.Next
		tmp = tmp.Next
	}
	tmpHead.Next = mid
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}

func reverse(head *ListNode) *ListNode {
	var front *ListNode
	for head != nil {
		next := head.Next
		head.Next = front
		if next == nil {
			break
		}
		front = head
		head = next
	}
	return head
}

func display(head *ListNode) {
	if head == nil {
		fmt.Println("----------")
		return
	}
	fmt.Println(head.Val)
	display(head.Next)
}

func main() {

}
