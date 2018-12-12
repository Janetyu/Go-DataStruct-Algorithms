package loopqueue

import (
	"reflect"
	"bytes"
	"fmt"
)

type LoopQueue struct {
	queueType     reflect.Kind
	data        []interface{}
	front, tail int
	size        int
}

func CreateLoopQueue(queueType reflect.Kind, cap int) *LoopQueue {
	// 因为循环队列会故意浪费一个空间，所以要在用户需要的空间上加1
	data := make([]interface{}, cap+1)
	size := 0
	return &LoopQueue{queueType, data,0,0, size}
}

// 无参数创建默认容量为10的数组，可自定义数组类型
func CreatedefaultLoopQueue(queueType reflect.Kind) *LoopQueue {
	return CreateLoopQueue(queueType,10)
}

func (queue *LoopQueue) GetSize() int {
	return queue.size
}

func (queue *LoopQueue) GetQueueCap() int {
	return cap(queue.data) - 1
}

func (queue *LoopQueue) IsEmpty() bool {
	return queue.front == queue.tail
}

func (queue *LoopQueue)Enqueue(e interface{}) {
	if (queue.tail + 1) % cap(queue.data) == queue.front {
		queue = resize(queue.GetQueueCap() * 2,queue)
	}

	queue.data[queue.tail] = e
	queue.tail = (queue.tail + 1) % cap(queue.data)
	queue.size ++
}

func (queue *LoopQueue)Dequeue() interface{} {
	if queue.IsEmpty() {
		panic("Cannot dequeue from an empty queue")
	}

	ret := queue.data[queue.front]
	queue.data[queue.front] = 0
	queue.front = (queue.front+1) % cap(queue.data)
	queue.size--

	if queue.size == queue.GetQueueCap() / 4 && queue.GetQueueCap() / 2 != 0 {
		queue = resize(queue.GetQueueCap() / 2,queue)
	}
	return ret
}

func (queue *LoopQueue)GetFront() interface{} {
	if queue.IsEmpty() {
		panic("Queue is empty.")
	}

	return queue.data[queue.front]
}

func resize(newcap int, queue *LoopQueue) *LoopQueue {
	newdata := make([]interface{}, newcap+1)
	for i := 0; i < queue.size; i++ {
		newdata[i] = queue.data[(i + queue.front) % cap(queue.data)]
	}

	queue.data = newdata
	queue.front = 0
	queue.tail = queue.size
	return queue
}

func (queue *LoopQueue) String() string {
	var buffer bytes.Buffer
	str1 := fmt.Sprintf("Queue: size = %d, capacity = %d \n", queue.GetSize(), queue.GetQueueCap())
	str2 := "front ["
	str3 := "] tail"
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	for i := queue.front; i != queue.tail; i = (i + 1) % cap(queue.data) {
		buffer.WriteString(fmt.Sprintf("%v", queue.data[i]))
		if (i + 1) % cap(queue.data) != queue.tail {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(str3)
	return buffer.String()
}
