package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ret := &ListNode{}
	tail := ret
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node := &ListNode{Val: l1.Val, Next: &ListNode{}}
			tail.Next = node
			l1 = l1.Next
		} else {
			node := &ListNode{Val: l2.Val, Next: &ListNode{}}
			tail.Next = node
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}
	return ret.Next
}
func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{}}
	l2 := &ListNode{Val: 2, Next: &ListNode{}}
	fmt.Println(mergeTwoLists(l1, l2))
}
