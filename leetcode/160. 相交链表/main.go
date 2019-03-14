package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		tmp  *ListNode = headA
		last *ListNode
	)
	if headA == nil || headB == nil {
		return nil
	}
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	last = tmp
	last.Next = headB

	slow, fast := headA, headA

	for {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			last.Next = nil
			return nil
		}
		fast = fast.Next
		if fast == nil {
			last.Next = nil
			return nil
		}

		if slow == fast {
			break
		}
	}
	slow = headA
	for {
		if slow == fast {
			last.Next = nil
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
}

func main() {

}
