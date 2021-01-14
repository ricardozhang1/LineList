package main

import (
	"fmt"
)

type QueueNode struct {
	Val  interface{}
	Next *QueueNode
	Prev *QueueNode
}

type QueueList struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func NewQueueNode(value interface{}) *QueueNode {
	return &QueueNode{Val: value}
}

func NewQueueList() *QueueList {
	return &QueueList{}
}

func (q *QueueList) IsEmpty() bool {
	if q.size == 0 {
		return true
	} else {
		return false
	}
}

func (q *QueueList) GetSize() int {
	return q.size
}

func (q *QueueList) DeQueue() *QueueNode {
	if q == nil {
		return nil
	} else {
		snode := q.tail
		q.tail = q.tail.Prev
		q.tail.Next = nil
		q.size--
		return snode
	}
}

func (q *QueueList) EnQueue(node *QueueNode) {
	if node == nil {
		return
	}

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		node.Next = q.head
		fmt.Println(node, q.head)
		q.head.Prev = node
		q.head = node
	}
	q.size++
}

func (q QueueList) PrintQueueList() {
	for q.head != nil {
		fmt.Println(q.head.Val)
		q.head = q.head.Next
	}
}

func main() {
	qq := NewQueueList()

	empty := qq.IsEmpty()
	fmt.Println(empty)

	qq.EnQueue(NewQueueNode("aa"))
	qq.EnQueue(NewQueueNode("bb"))
	qq.EnQueue(NewQueueNode("cc"))
	qq.EnQueue(NewQueueNode("dd"))
	qq.EnQueue(NewQueueNode("ee"))
	qq.EnQueue(NewQueueNode("ff"))
	qq.EnQueue(NewQueueNode("gg"))
	qq.EnQueue(NewQueueNode("hh"))

	qSize := qq.GetSize()
	fmt.Println(qSize)

	qq.PrintQueueList()
	empty = qq.IsEmpty()
	fmt.Println(empty)

	var snode *QueueNode
	snode = qq.DeQueue()
	fmt.Println("1", snode.Val)
	snode = qq.DeQueue()
	fmt.Println("2", snode.Val)
	snode = qq.DeQueue()
	fmt.Println("3", snode.Val)
	snode = qq.DeQueue()
	fmt.Println("4", snode.Val)
	qSize = qq.GetSize()
	fmt.Println(qSize)

	qq.PrintQueueList()
}
