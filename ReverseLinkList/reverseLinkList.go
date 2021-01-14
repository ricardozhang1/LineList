package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{Val: value, Next: nil}
}

type LinkList struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewLinkList() *LinkList {
	return &LinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (nl *LinkList) InsertByTail(value int) {
	node := NewListNode(value)

	if nl.size == 0 {
		nl.head = node
		nl.tail = node
	} else {
		nl.tail.Next = node
		nl.tail = node
	}
	nl.size++
}

func (nl LinkList) Print() {
	if nl.size == 0 {
		return
	}
	node := nl.head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

// reverseList 链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	p.Next = head
	head.Next = nil
	return head
}

// reverseList2 链表反转
func reverseList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var prev *ListNode
    cur := head
    for cur != nil {
        cur.Next, prev, cur = prev, cur, cur.Next
    }
    return prev
}

// 快慢指针查找中位数
func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	ls := NewLinkList()

	for i := 0; i < 20; i++ {
		ls.InsertByTail(i + 2)
	}

	ls.Print()

	fmt.Println("=========================")
	_ = reverseList(ls.head)
	// for newList != nil {
	// 	fmt.Println(newList.Val)
	// 	newList = newList.Next
	// }

	temp := ls.head
	ls.head = ls.tail
	ls.tail = temp
	fmt.Println(ls.head, ls.tail)
	ls.Print()

	fmt.Println("=========================")
	d := getMid(ls.head)
	fmt.Println(d.Val)

	// fmt.Println("=========================")
	// newList2 := reverseList2(ls.head)
	// for newList2 != nil {
	// 	fmt.Println(newList2.Val)
	// 	newList2 = newList2.Next
	// }
}
