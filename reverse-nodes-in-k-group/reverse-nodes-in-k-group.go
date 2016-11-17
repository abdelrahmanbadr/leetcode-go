package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	ret := &ListNode{}
	tail := ret
	cursor := head
	lists := make([]*ListNode, k)

	count := 0
	for cursor != nil {
		lists[count] = cursor
		count++
		cursor = cursor.Next
		if count == k {
			for i := k - 1; i > 0; i-- {
				lists[i].Next = lists[i-1]
			}
			tail.Next = lists[k-1]
			tail = lists[0]
			lists[0] = nil
			count = 0
		}
	}
	tail.Next = lists[0]
	return ret.Next

}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	r := reverseKGroup(l1, 2)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
