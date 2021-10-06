package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := new(ListNode)
	ptmp := newListNode

	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			ptmp.Next = l1
			l1 = l1.Next

		} else {
			ptmp.Next = l2
			l2 = l2.Next
		}
		ptmp = ptmp.Next
	}

	return newListNode.Next
}

func main() {
	//初始化链表
	node1 := new(ListNode)
	node1.Val = 1

	node11 := new(ListNode)
	node11.Val = 2
	node1.Next = node11

	node111 := new(ListNode)
	node111.Val = 4
	node11.Next = node111

	node2 := new(ListNode)
	node2.Val = 1

	node22 := new(ListNode)
	node22.Val = 3
	node2.Next = node22

	node222 := new(ListNode)
	node222.Val = 4
	node22.Next = node222

	var newListNode *ListNode
	newListNode = mergeTwoLists(node1, node2)
	for newListNode != nil {
		fmt.Println(*newListNode)
		newListNode = newListNode.Next
	}
}

