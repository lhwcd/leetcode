package mergetwolists

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	cur := head
	//
	for l1 != nil {
		if l2 != nil {
			//
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else {
			cur.Next = l1
			return head
		}
	}
	//
	if l2 != nil {
		cur.Next = l2
	}
	//
	return head
}
