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

func hasCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var (
		fast    *ListNode = head.Next.Next
		slow    *ListNode = head.Next
		newNode           = head
	)
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	for newNode != slow {
		newNode = newNode.Next
		slow = slow.Next
	}

	return newNode
}

func main() {
}

// 我认为这道题更像是数学题

// 假设有环链表，设节点A为入环节点(也就是我们想要的节点)，节点B为快慢指针相遇的节点。
// 设头节点到节点A的距离为x，节点A到节点B的距离为y。
// 在快慢指针相遇时，慢指针走了x+y的距离，而快指针走的距离是慢指针的两倍，所以快指针走的
// 距离是2(x+y)。由此可知，慢指针在走x+y必然可以回到节点B(理解这个非常重要)。
// 重点来了，假如我们的慢指针在向前走x距离，那它到哪了？是不是到了节点A。(因为在往前走y距离就又回到了节点B)
// 当慢指针再次走的时候，让一个新的指针从头节点出发走x距离，它就会到达节点A。这时慢指针和新指针在节点A相遇。
