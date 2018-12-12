package queue

import (
	"reflect"
	"testing"
)

func Test_Queue(t *testing.T) {
	var queue Queue
	queue = CreatedefaultQueue(reflect.Int)

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		t.Log(queue)

		if i%3 == 2 {
			queue.Dequeue()
			t.Log(queue)
		}
	}

}
