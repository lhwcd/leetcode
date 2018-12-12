package mergetwolists

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

func newList(start, end int) *ListNode {
	head := &ListNode{
		Val: start,
	}
	cur := head
	for i := start + 1; i < end; i++ {
		cur.Next = &ListNode{
			Val: i,
		}
		cur = cur.Next
	}
	return head
}
func TestMergeTwoLists(t *testing.T) {
	list1 := newList(1, 5)
	printList(list1)
	list2 := newList(4, 9)
	printList(list2)
	list := mergeTwoLists(list1, list2)
	printList(list)
}
