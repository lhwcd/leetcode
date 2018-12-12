package checklistcircle

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//checkListCircle 如果链表有环，返回true，否则false
func checkListCircle(list *ListNode) bool {
	if list == nil || list.Next == nil {
		return false
	}
	//
	slow, fast := list, list
	//
	for fast != nil && fast.Next != nil {
		//
		fast = fast.Next.Next
		slow = slow.Next
		//
		if fast == slow {
			return true
		}
	}
	return false
}
