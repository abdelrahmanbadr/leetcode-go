package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func minVal(lists []*ListNode) (bool, int) {
	has, min, min_node := false, 2147483647, 0
	for idx, l := range lists {
		if l != nil {
			has = true
			if l.Val <= min {
				min = l.Val
				min_node = idx
			}
		}
	}
	if has {
		lists[min_node] = lists[min_node].Next
	}
	return has, min
}

func mergeKLists(lists []*ListNode) *ListNode {
	ret := &ListNode{}
	tail := ret
	for {
		has, min := minVal(lists)
		if has == false {
			return ret.Next
		} else {
			node := &ListNode{Val: min, Next: nil}
			tail.Next = node
			tail = tail.Next
		}
	}
	return ret.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: nil}
	l2 := &ListNode{Val: 2, Next: nil}
	l3 := &ListNode{Val: 3, Next: nil}
	fmt.Println(mergeKLists([]*ListNode{l1, l2, l3}))
}
