package main

// 给定一个链表，判断链表中是否有环。

// 进阶：
// 你能否不使用额外空间解决此题？

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(_head *ListNode) bool {
	if _head == nil {
		return false
	}
	for _head.Next != nil {
		if _head == _head.Next {
			return true
		}
		_head.Next = _head.Next.Next
	}
	return false
}

func main() {

}
