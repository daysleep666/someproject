package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	var ji, jihead, ou, ouhead *ListNode
	for i := 1; ; i++ {
		if head == nil {
			break
		}
		next := head.Next
		if i%2 != 0 { // 奇数
			if ji == nil {
				ji = head
				jihead = head
			} else {
				ji.Next = head
				ji = ji.Next
			}
		} else { //偶数
			if ou == nil {
				ou = head
				ouhead = head
			} else {
				ou.Next = head
				ou = ou.Next
			}
		}
		head = next
	}
	if ji != nil {
		ji.Next = ouhead
	}
	if ou != nil {
		ou.Next = nil
	}
	return jihead
}

func main() {
	oddEvenList(nil)
}
