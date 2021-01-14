package main

import (
	"fmt"
)

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{Val: value}
}

func (node *ListNode) PrintList() {
	for node!=nil && node.Next!=nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

// getEnter 获取环的入口节点，并返回
func (node *ListNode) getEnter() *ListNode {
	fast := node
	slow := node
	var temp *ListNode
	for fast!=nil && fast.Next!=nil && fast.Next.Next!=nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast.Val.(string) == slow.Val.(string) {
			temp = node
			continue
		}

		if temp != nil {
			temp = temp.Next
			if temp.Val.(string) == slow.Val.(string) {
				return temp
			}
		}
	}
	return nil

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
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = sixth
	sixth.Next = seventh

	first.PrintList()

	//产生环
	seventh.Next = third

	entrance := first.getEnter()
	fmt.Println("first链表的环的入口节点是：", entrance)
}
