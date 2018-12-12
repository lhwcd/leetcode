package reverselist

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

func TestReverseList(t *testing.T) {

	list := newList(1)
	printList(list)
	list = reverseList(list)
	printList(list)

	list = newList(2)
	printList(list)
	list = reverseList(list)
	printList(list)

	list = newList(5)
	printList(list)
	list = reverseList(list)
	printList(list)
}
