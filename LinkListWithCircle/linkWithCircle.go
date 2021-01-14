package main

import (
	"fmt"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{Val: value}
}

type ListLink struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewLinkList() *ListLink {
	return &ListLink{}
}

func (l *ListLink) InsertByTail(node *ListNode) {
	if node == nil {
		return
	}

	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}

	l.size++
}

func (l ListLink) PrintLink() {
	if l.size == 0 {
		return
	} else {
		for l.head != nil {
			fmt.Println(l.head.Val)
			l.head = l.head.Next
		}
	}
}

func (l ListLink) isCircle() bool {
	slow := l.head
	fast := l.head
	for fast.Next.Next!=nil {	
		fast = fast.Next.Next
		slow = slow.Next

		if fast.Val.(int) == slow.Val.(int) {
			return true
		}
	}
	return false
}

func main() {
	ls := NewLinkList()

	for i := 1; i < 7; i++ {
		snode := NewListNode(i)
		ls.InsertByTail(snode)
	}

	ls.PrintLink()
	//不存在环 返回false
	d := ls.isCircle()
	fmt.Println("是否存在环：", d)

	fmt.Println(ls.head, ls.tail)
	// 形成一个环 1->2->3->4->5->6 || 6->3的环
	ls.tail.Next = ls.head.Next.Next
	// ls.PrintLink()  // 存在环，会出现无限循环
	
	// 存在环，返回true
	dd := ls.isCircle()
	fmt.Println("是否存在环：", dd)

}
