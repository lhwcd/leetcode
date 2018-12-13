package listmiddlenode

import (
	"fmt"
	"testing"
)

func printList(list *ListNode) {
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Println()
}

func newList(n int) *ListNode {
	if n == 0 {
		return nil
	}
	//
	head := &ListNode{
		Val: 0,
	}
	cur := head
	//
	for i := 1; i < n; i++ {
		cur.Next = &ListNode{
			Val: i,
		}
		cur = cur.Next
	}
	return head
}

func TestMiddleNode(t *testing.T) {
	list := newList(6)
	printList(list)
	list = middleNode(list)
	printList(list)
}
