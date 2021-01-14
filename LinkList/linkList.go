package main

import (
	"fmt"
)

type LinkNode struct {
	Data interface{}
	next *LinkNode
}

type SingleLinkList struct {
	size int
	head *LinkNode
	tail *LinkNode
}

//初始化链表
func InitSingleLink() *SingleLinkList {
	return &SingleLinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

//获取头部节点
func (sl *SingleLinkList) GetHead() *LinkNode {
	return sl.head
}

//获取尾部节点
func (sl *SingleLinkList) GetTail() *LinkNode {
	return sl.tail
}

//打印链表
func (sl SingleLinkList) Print() {
	fmt.Println("SingleLinkList size: ", sl.Length())
	if sl.size == 0 {
		return
	}

	ptr := sl.GetHead()
	for ptr != nil {
		fmt.Println("Data: ", ptr.Data)
		ptr = ptr.next
	}
}

//链表长度
func (sl *SingleLinkList) Length() int {
	return sl.size
}

//查询数据
func (sl *SingleLinkList) Search(index int) *LinkNode {
	if index < 0 || index > sl.Length() || sl.Length() == 0 {
		return nil
	}

	// 是否是头部节点
	if index == 0 {
		return sl.GetHead()
	}

	node := sl.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

//头部插入数据
func (sl *SingleLinkList) InsertByHead(node *LinkNode) {
	if node == nil {
		return
	}

	//判断是否第一个节点
	if sl.Length() == 0 {
		sl.head = node
		sl.tail = node
		node.next = nil
	} else {
		oldHeadNode := sl.GetHead()
		sl.head = node
		sl.head.next = oldHeadNode
	}
	sl.size++
}

//尾部插入数据
func (sl *SingleLinkList) InsertByTail(node *LinkNode) {
	if node == nil {
		return
	}

	//插入第一个节点
	if sl.size == 0 {
		sl.head = node
		sl.tail = node
		node.next = nil
	} else {
		sl.tail.next = node
		node.next = nil
		sl.tail = node
	}
	sl.size++
}

//插入数据，中间位置
func (sl *SingleLinkList) InsertByIndex(index int, node *LinkNode) {
	if node == nil || index < 0 {
		return
	}

	//往头部插入
	if index == 0 {
		sl.InsertByHead(node)
	} else {
		if index > sl.Length() {
			return
		} else if index == sl.Length() {
			// 往尾部添加
			sl.InsertByTail(node)
		} else {
			preNode := sl.Search(index - 1)
			currentNode := sl.Search(index)
			preNode.next = node
			node.next = currentNode
			sl.size++
		}
	}
}

//删除数据，根据下标
func (sl *SingleLinkList) DeleteByIndex(index int) {
	if index < 0 || index > sl.Length() || sl.Length() == 0 {
		return
	}

	// 删除第一个节点
	if index == 0 {
		sl.head = sl.head.next
	} else {
		preNode := sl.Search(index - 1)
		if index != sl.Length()-1 {
			node := sl.Search(index)
			preNode.next = node.next
		} else {
			//删除尾结点
			sl.tail = preNode
			preNode.next = nil
		}
	}
	sl.size--
}

//删除数据 根据数据 ??
func (sl *SingleLinkList) DeleteByData(data interface{}) {
	if sl.Length() == 0 || data == nil {
		return
	}

	node := sl.head
	preNode := sl.head
	for node.next != nil {
		preNode = node
		node = node.next
		if node.Data.(int) == data.(int) {
			preNode.next = node.next
			node.next = nil
			node.Data = nil
			node = nil
			return
		}
	}
}

//销毁链表
func (sl *SingleLinkList) Destory() {
	sl.tail = nil
	sl.head = nil
	sl.size = 0
}

func main() {
	// 初始化链表
	sl := InitSingleLink()

	// 指定指标插入
	for i := 0; i < 5; i++ {
		snode := &LinkNode{Data: i, next: nil}
		sl.InsertByIndex(i, snode)
	}

	sl.Print()
	fmt.Println("============================================")

	var snode *LinkNode

	//往头部插入节点
	snode = &LinkNode{Data: 100, next: nil}
	sl.InsertByHead(snode)
	sl.Print()
	fmt.Println("============================================")

	//往尾部插入节点
	snode = &LinkNode{Data: 200, next: nil}
	sl.InsertByTail(snode)
	sl.Print()
	fmt.Println("============================================")

	// 查询下标为2的节点
	node := sl.Search(2)
	fmt.Println("Node2: ", node.Data)
	sl.Print()
	fmt.Println("============================================")

	//删除下标为2的节点
	sl.DeleteByIndex(2)
	sl.Print()
	fmt.Println("============================================")

	//删除Data为3的节点
	sl.DeleteByData(3)
	sl.Print()
}
