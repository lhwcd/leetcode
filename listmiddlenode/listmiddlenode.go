package listmiddlenode

/**
 * Definition for singly-linked list.
 */

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//
func middleNode(head *ListNode) *ListNode {
	//
	if head == nil || head.Next == nil {
		return head
	}
	//
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
