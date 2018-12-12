package loopqueue

import (
	"testing"
	"reflect"

	. "Go-DataStruct-Algorithms/Data-structures/queue"
)

func Test_LoopQueue(t *testing.T) {
	var queue Queue
	queue = CreatedefaultLoopQueue(reflect.Int)

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		t.Log(queue)

		if i%3 == 2 {
			queue.Dequeue()
			t.Log(queue)
		}
	}

}
