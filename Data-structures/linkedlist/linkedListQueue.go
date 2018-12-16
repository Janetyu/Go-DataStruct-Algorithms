package linkedlist

import "bytes"

type LinkedListQueue struct {
	head,tail *Node // 定义头结点和尾结点
	size int
}

func CreateLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{size: 0}
}

func (queue *LinkedListQueue)GetSize() int {
	return queue.size
}

func (queue *LinkedListQueue)IsEmpty() bool {
	return queue.size == 0
}

func (queue *LinkedListQueue)Enqueue(e interface{}) {
	if queue.tail == nil {
		queue.tail = &Node{e: e,next: nil}
		queue.head = queue.tail
	} else {
		queue.tail.next = &Node{e: e,next: nil}
		queue.tail = queue.tail.next
	}
	queue.size++
}

func (queue *LinkedListQueue)Dequeue() interface{} {
	if queue.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}

	retNode := queue.head
	queue.head = queue.head.next
	retNode.next = nil
	if queue.head == nil {
		queue.tail = nil
	}
	queue.size--

	return retNode.e
}

func (queue *LinkedListQueue)GetFront() interface{} {
	if queue.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	return queue.head.e
}

func (queue *LinkedListQueue)String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Queue: front ")

	cur := queue.head
	for cur != nil {
		buffer.WriteString(cur.String() + " -> ")
		cur = cur.next
	}

	buffer.WriteString("NULL tail")
	return buffer.String()
}