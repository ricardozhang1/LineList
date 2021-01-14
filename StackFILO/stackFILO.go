package main

import (
	"fmt"
)

type StackNode struct {
	Vale interface{}
	Next *StackNode
}

func NewStackNode(value interface{}) *StackNode {
	return &StackNode{Vale: value}
}

type StackLink struct {
	head *StackNode
	size int
}

func NewStackLink() *StackLink {
	return &StackLink{head: nil, size: 0}
}

func (sl *StackLink) IsEmpty() bool {
	if sl.size == 0 {
		return true
	} else {
		return false
	}
}

func (sl *StackLink) Size() int {
	return sl.size
}

func (ls *StackLink) Pop() *StackNode {
	if ls.head == nil {
		return nil
	} else {
		node := ls.head
		ls.head = ls.head.Next
		return node
	}
}

func (ls *StackLink) Push(value interface{}) {
	if value == nil {
		return
	}

	snode := NewStackNode(value)
	if ls.head == nil {
		ls.head = snode
	} else {
		snode.Next = ls.head
		ls.head = snode
	}
}

func (ls StackLink) PrintStackLink() {
	for ls.head != nil {
		fmt.Println(ls.head.Vale)
		ls.head = ls.head.Next
	}
}

func main() {
	// sl := NewStackLink()
	// sl.Push("aa")
	// sl.Push("bb")
	// sl.Push("cc")
	// sl.Push("dd")
	// sl.Push("ee")
	// sl.Push("ff")
	// sl.Push("gg")
	// length := sl.Size()
	// sl.PrintStackLink()

	// fmt.Println(length)

	// node := sl.Pop()
	// fmt.Println(node)
	// node = sl.Pop()
	// fmt.Println(node)
	// node = sl.Pop()
	// fmt.Println(node)
	// length = sl.Size()
	// sl.PrintStackLink()
	

	// 括号匹配问题
	ssl := NewStackLink()
	
	strWord := "(上海)(长((长安)安))"
	for _, v := range []rune(strWord) {
		if string(v) == "(" {
			ssl.Push(string(v))
		} else if string(v) == ")" {
			v := ssl.Pop()
			if v == nil {
				// fmt.Println("括号不匹配！")
				ssl.head = NewStackNode(v)
			}
		}
	}

	if ssl.head != nil {
		fmt.Println("括号不匹配！")
	} else {
		fmt.Println("括号匹配！")
	}
	

}
