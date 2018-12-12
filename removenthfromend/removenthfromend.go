package removenthfromend

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//
	if head == nil || n == 0 {
		return head
	}
	//
	fast, slow := head, head
	for i := 0; i < n; i++ {
		//注意的是输入的链表只有一个节点，要求删除一个节点的情况，也就是n=1的情况
		if fast.Next == nil {
			return head.Next
		}
		fast = fast.Next
	}
	//
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//
	slow.Next = slow.Next.Next
	return head
}
