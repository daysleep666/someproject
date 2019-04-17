package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var front *ListNode
	for fast != nil {
		fast = fast.Next
		if fast == nil { // 说明是奇数个节点
			slow = slow.Next
			break
		}
		fast = fast.Next

		tmpNext := slow.Next
		slow.Next = front
		front = slow
		slow = tmpNext
		tmpNext = tmpNext.Next
	}
	for slow != nil && front != nil && slow.Val == front.Val {
		slow = slow.Next
		front = front.Next
	}
	return slow == nil && front == nil
}

func main() {

}
