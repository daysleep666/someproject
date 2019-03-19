package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	num := 1
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		num++
	}
	k = k % num
	tmp.Next = head

	tmp = head
	for i := 1; i < num-k; i++ {
		tmp = tmp.Next
	}
	head = tmp.Next
	tmp.Next = nil
	return head
}

func main() {

}
