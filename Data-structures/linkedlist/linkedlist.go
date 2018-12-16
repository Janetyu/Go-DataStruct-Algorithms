package linkedlist

import (
	"fmt"
	"bytes"
)

type LinkedList struct {
	dummyHead *Node // 虚拟头结点
	size int
}

type Node struct {
	e interface{}
	next *Node
}

func CreateNode(e interface{},next *Node) *Node {
	return &Node{e,next}
}

func CreateDataNode(e interface{}) *Node {
	return CreateNode(e,&Node{})
}

func CreateNilNode() *Node {
	return CreateNode(nil,&Node{})
}

func (n *Node)String() string {
	return fmt.Sprint(n.e)
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{dummyHead: &Node{next:nil,e:nil},size: 0}
}

// 获取链表中元素个数
func (list *LinkedList)GetSize() int {
	return list.size
}

// 返回链表是否为空
func (list *LinkedList)IsEmpty() bool {
	return list.size == 0
}

// 向链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用:)
func (list *LinkedList)Add(index int, e interface{}) {

	if index < 0 || index > list.size {
		panic("Add failed.Illegal index.")
	}

	// 这是有头结点且不为虚拟结点时的逻辑
	//if index == 0 {
	//	list.AddFirst(e)
	//} else {
	//	prev := list.head
	//	for i := 0; i < index - 1; i++ {
	//		prev = prev.next
	//	}

	//	//n := Node{e,Node{}}
	//	//n.next = prev.next
	//	//prev.next = n

	//	prev.next = Node{e, prev.next}
	//	list.size++
	//}

	// 这是虚拟头结点的逻辑，注意遍历到index之前
	// 注意遍历的目的，如果是获取index的前一个元素，则prev应该为dummyHead，从头结点开始遍历
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{e, prev.next}
	list.size++
}

// 向链表头添加新的元素e
func (list *LinkedList)AddFirst(e interface{}) {
	//n := Node{e,Node{}}
	//n.next = list.head
	//list.head = n

	// 上面三句话等同于下面这句
	//list.head = Node{e, list.head}
	//
	//list.size++

	list.Add(0, e)
}

// 在链表末尾添加新的元素e
func (list *LinkedList)AddLast(e interface{})  {
	list.Add(list.size,e)
}

// 获得链表的第index(0-based)位置的元素
// 在链表中不是一个常用的操作，练习用:)
func (list *LinkedList)Get(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Get failed.Illegal index.")
	}

	// 注意遍历的目的，如果是获取index的元素，则cur应该为dummyHead.next，从第一个元素开始遍历
	cur := list.dummyHead.next
	for i:=0; i < index ; i++ {
		cur = cur.next
	}
	return cur.e
}

// 获取链表的第一个元素
func (list *LinkedList)GetFirst() interface{} {
	return list.Get(0)
}

// 获取链表的最后一个元素
func (list *LinkedList)GetLast() interface{} {
	return list.Get(list.size - 1)
}

// 修改链表的第index(0-based)位置的元素
// 在链表中不是一个常用的操作，练习用:)
func (list *LinkedList)Set(index int, e interface{}) {
	if index < 0 || index >= list.size {
		panic("Set failed.Illegal index.")
	}

	cur := list.dummyHead.next
	for i:=0; i < index ; i++ {
		cur = cur.next
	}
	cur.e = e
}

// 查找链表中是否有元素e
func (list *LinkedList)Contains(e interface{}) bool {

	cur := list.dummyHead.next
	for cur.e != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}

	return false
}

// 删除链表的第index(0-based)位置的元素
// 在链表中不是一个常用的操作，练习用:)
func (list *LinkedList)Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Remove failed.Illegal index.")
	}

	// 注意遍历的目的，如果是获取index的元素，则cur应该为dummyHead.next，从第一个元素开始遍历
	prev := list.dummyHead
	for i:=0; i < index ; i++ {
		prev = prev.next
	}

	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil

	list.size--
	return retNode.e
}

// 删除链表的第一个元素并返回值
func (list *LinkedList)RemoveFirst() interface{} {
	return list.Remove(0)
}

// 删除链表的最后一个元素并返回值
func (list *LinkedList)RemoveLast() interface{} {
	return list.Remove(list.size - 1)
}

func (list *LinkedList)String() string {
	var buffer bytes.Buffer
	str := "NULL"

	//cur := list.dummyHead.next
	//for cur.e != nil {
	//	buffer.WriteString(cur.String() + " -> ")
	//	cur = cur.next
	//}

	for cur := list.dummyHead.next ; cur != nil ; cur = cur.next {
		buffer.WriteString(cur.String() + " -> ")
	}
	buffer.WriteString(str)

	return buffer.String()
}



