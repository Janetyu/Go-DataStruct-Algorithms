package queue

import (
	"Go-DataStruct-Algorithms/Data-structures/array"
	"reflect"
	"bytes"
	"fmt"
)

type ArrayQueue struct {
	array *array.Array
}

func CreateQueue(queueType reflect.Kind, cap int) *ArrayQueue {
	return &ArrayQueue{array.CreateArray(queueType,cap)}
}

func CreatedefaultQueue(queueType reflect.Kind) *ArrayQueue {
	return &ArrayQueue{array.CreatedefaultArray(queueType)}
}

func (queue *ArrayQueue)GetSize() int {
	return queue.array.GetArrLen()
}

func (queue *ArrayQueue)IsEmpty() bool {
	return queue.array.IsEmpty()
}

func (queue *ArrayQueue)GetCapacity() int {
	return queue.array.GetArrCap()
}

func (queue *ArrayQueue)Enqueue(e interface{})  {
	queue.array.AddLast(e)
}

func (queue *ArrayQueue)Dequeue() interface{} {
	return queue.array.RemoveFirst()
}

func (queue *ArrayQueue)GetFront() interface{} {
	return queue.array.GetFirst()
}

func (queue *ArrayQueue) String() string {
	var buffer bytes.Buffer
	str1 := fmt.Sprintf("Queue: size = %d, capacity = %d \n", queue.array.GetArrLen(), queue.array.GetArrCap())
	str2 := "front ["
	str3 := "] tail"
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	for i := 0; i < queue.array.GetArrLen(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", queue.array.Get(i)))
		if i != queue.array.GetArrLen()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(str3)
	return buffer.String()
}
