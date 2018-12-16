package linkedlist

import "bytes"

type LinkedListStack struct {
	linkedlist *LinkedList
}

func CreateLinkedListStack() *LinkedListStack {
	return &LinkedListStack{CreateLinkedList()}
}

func (list *LinkedListStack)GetSize() int {
	return list.GetSize()
}

func (list *LinkedListStack)IsEmpty() bool {
	return list.IsEmpty()
}

func (list *LinkedListStack)Push(e interface{})  {
	list.linkedlist.AddFirst(e)
}

func (list *LinkedListStack)Pop() interface{} {
	return list.linkedlist.RemoveFirst()
}

func (list *LinkedListStack)Peek() interface{} {
	return list.linkedlist.GetFirst()
}

func (list *LinkedListStack)String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Stack: top ")
	buffer.WriteString(list.linkedlist.String())
	return buffer.String()
}
