package checklistcircle

import (
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func printList(list *ListNode) {
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Println()
}

//
func newNoCirclieList(n int) *ListNode {
	head := &ListNode{
		Val: 0,
	}
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = &ListNode{
			Val: i,
		}
		cur = cur.Next
	}
	return head

}

//
func newCirclieList(n int) *ListNode {
	//
	head := &ListNode{
		Val: 0,
	}
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = &ListNode{
			Val: i,
		}
		cur = cur.Next
	}
	//
	cur.Next = head.Next.Next
	//
	return head
}

//
func TestCheckListCircle(t *testing.T) {
	//
	list := newNoCirclieList(5)
	printList(list)
	//
	expect := false
	output := checkListCircle(list)
	assert.Equal(t, expect, output)

	//
	list = newCirclieList(6)
	expect = true
	output = checkListCircle(list)
	assert.Equal(t, expect, output)
}
