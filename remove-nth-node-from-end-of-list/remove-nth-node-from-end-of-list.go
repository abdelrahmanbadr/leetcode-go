package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	tempNode := head
	count := 1
	for tempNode.Next != nil {
		count++
		tempNode = tempNode.Next
	}
	if count == n {
		return head.Next
	}
	tempNode = head
	for i := 0; i < count-n-1; i++ {
		tempNode = tempNode.Next
	}
	tempNode.Next = tempNode.Next.Next
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(removeNthFromEnd(head, 2))
}
