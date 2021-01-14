package main

import (
	"fmt"
)

type ListNode struct {
	Val interface{}
	Prev *ListNode
	Next *ListNode
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{Val: value}
}

func (node *ListNode) PrintList() {
	if node==nil {
		return
	}

	for node!=nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	first := NewListNode("aa")
	second := NewListNode("bb")
	third := NewListNode("cc")
	fourth := NewListNode("dd")
	fifth := NewListNode("ee")
	sixth := NewListNode("ff")
	seventh := NewListNode("gg")

	first.Next = second
	second.Prev = first
	second.Next = third
	third.Prev = second
	third.Next = fourth
	fourth.Prev = third
	fourth.Next = fifth
	fifth.Prev = fourth
	fifth.Next = sixth
	sixth.Prev = fifth
	sixth.Next = seventh
	seventh.Prev = sixth

	first.PrintList()

	//构建循环链表，让最后一个节点指向第一个节点
	first.Prev = seventh
	seventh.Next = first
	// 使用count，记录当前的报数值
	var count int

	n := first
	var before *ListNode
	for (n!=n.Next) {
		count++
		if count==3 {
			before.Next = n.Next
			fmt.Println(n.Val.(string) + "==>")
			count = 0
			n = n.Next
		} else {
			before = n
			n = n.Next
		}
	}
	fmt.Println("the last one: ", n.Val)
}


