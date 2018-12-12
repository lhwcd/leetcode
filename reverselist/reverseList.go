package reverselist

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//reverseList 输入一个单链表，返回一个反转后的链表
func reverseList(list *ListNode) *ListNode {
	//
	if list == nil {
		return list
	}
	//
	var pre *ListNode
	var next *ListNode
	var cur = list
	for cur != nil {
		next = cur.Next //save next
		cur.Next = pre  //
		pre = cur       //move next
		cur = next      //move next
	}
	return pre
}
