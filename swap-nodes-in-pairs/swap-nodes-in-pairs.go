package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	tail := &ListNode{Val: 0, Next: head}
	ret := tail
	for tail.Next != nil && tail.Next.Next != nil {
		node1 := tail.Next
		node2 := tail.Next.Next
		node1.Next = node2.Next
		node2.Next = node1
		tail.Next = node2
		tail = node1
	}
	return ret.Next

}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println(swapPairs(l1))
}
