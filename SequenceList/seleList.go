package main

import (
	"fmt"
)

//数据结构之线性表--顺序表
type List struct {
	Len      int            // 线性表的长度
	Capacity int            // 线性表的容量
	Prt      *[]interface{} // 指向线性表空间指针
}

// 初始化
func (l *List) ListInit(capacity int) {
	l.Capacity = capacity
	l.Len = 0
	m := make([]interface{}, capacity)
	l.Prt = &m
}

// is empty
func (l *List) ListIsEmpty() bool {
	if l.Len == 0 {
		return true
	} else {
		return false
	}
}

// is ful
func (l *List) ListIsFul() bool {
	if l.Len == l.Capacity {
		return true
	} else {
		return false
	}
}

//根据下标Get元素
func (l *List) ListGet(index int) (interface{}, bool) {
	if index < 0 || index > l.Len {
		return nil, false
	} else {
		return (*l.Prt)[index], true
	}
}

// 根据传入的值，返回匹配的元素下标，元素第一次出现的位置
func (l *List) ListLocal(elem interface{}) (int, bool) {
	for i, _ := range *l.Prt {
		if elem == (*l.Prt)[i] {
			return i, true
		}
	}
	return -1, false
}

// 寻找元素的前驱(当前元素的前一个元素)
func (l *List) ListElemPre(elem interface{}) (interface{}, bool) {
	i, _ := l.ListLocal(elem)
	// 顺序表中不存在该元素，或者元素为第一个元素，无前驱元素
	if i == -1 || i == 0 {
		return nil, false
	} else {
		pre := (*l.Prt)[i-1]
		return pre, true
	}
}

// 寻找元素的后驱元素
func (l *List) ListElemNext(elem interface{}) (interface{}, bool) {
	i, _ := l.ListLocal(elem)
	// 顺序表中不存在该元素，或者该元素为最后一个元素，无后驱元素
	if i == -1 || i == l.Len-1 {
		return nil, false
	} else {
		N := (*l.Prt)[i+1]
		return N, true
	}
}

// 插入元素，index为插入的位置，elem为插入的值
func (l *List) ListInsert(index int, elem interface{}) bool {
	//判断下标是否有效，以及表是否满
	if l.ListIsFul() {
		l.ListResize(l.Len*2)
	}

	if index < 0 || index > l.Capacity {
		return false
	} else {
		// 先将index位置元素以及之后元素后移一位
		for i := l.Len - 1; i >= index; i-- {
			(*l.Prt)[i+1] = (*l.Prt)[i]
		}
		// 插入元素
		(*l.Prt)[index] = elem
		l.Len++
		return true
	}
	
}

// 删除元素
func (l *List) ListDelete(index int) bool {
	//判断下标的有效性，以及表是否为空
	if index < 0 || index > l.Capacity || l.ListIsEmpty() {
		return false
	} else {
		//注意边界
		for i := index; i < l.Len-1; i++ {
			(*l.Prt)[i] = (*l.Prt)[i+1]
		}
		l.Len--

		if l.Len > 0 && l.Len < l.Capacity/4 {
			l.ListResize(l.Capacity/2)
		}
		return true
	}


}

//遍历
func (l *List) ListTraverse() {
	for i := 0; i < l.Len; i++ {
		fmt.Println((*l.Prt)[i])
	}
}

//清空
func (l *List) ListClear() {
	l.Len = 0
	// 指针为空
	l.Prt = nil
}

//改变容量
func (l *List) ListResize(newcapacity int) {
	// 记录旧数组
	var temp *[]interface{} = l.Prt
	// 创建新数组
	m := make([]interface{}, newcapacity)
	l.Prt = &m
	for i:=0; i<l.Len; i++ {
		(*l.Prt)[i] = (*temp)[i]
	}
	l.Capacity = newcapacity
}

func main() {
	var li List
	li.ListInit(2)
	fmt.Println(li.ListIsEmpty())
	fmt.Println(li.ListIsFul())

	type p struct {
		name string
		number int
	}

	elem1 := p{"yaoming", 1}
	elem2 := p{"kebi", 4}
	elem3 := p{"maidi", 9}
	elem4 := p{"zhanmusi", 6}
	elem5 := p{"aaaa", 5}
	elem6 := p{"cccc", 12}

	li.ListInsert(0, elem1)

	li.ListTraverse()

	li.ListInsert(1, elem2)
	li.ListInsert(2, elem3)

	li.ListTraverse()

	li.ListInsert(3, elem4)
	li.ListInsert(4, elem5)
	li.ListInsert(5, elem6)

	//测试获取
	// ret, ok := li.ListGet(1)
	// if !ok {
	// 	fmt.Println("获取元素失败！")
	// }
	// fmt.Println("GET index 1: ", ret)

	//测试删除
	// li.ListDelete(2)

	//遍历测试
	li.ListTraverse()

	//重新确定数据大小
	// li.ListResize(10)
	fmt.Println(li.ListIsEmpty())
	fmt.Println(li.ListIsFul())
	fmt.Println(li.Capacity, li.Len)

	li.ListDelete(1)
	li.ListDelete(1)
	li.ListDelete(1)
	li.ListDelete(1)
	li.ListDelete(1)
	li.ListDelete(1)
	fmt.Println(li.ListIsEmpty())
	fmt.Println(li.ListIsFul())
	fmt.Println(li.Capacity, li.Len)


}
