package Data_structures

import (
	"testing"
	"reflect"
	"time"

	. "Go-DataStruct-Algorithms/Data-structures/queue"
	. "Go-DataStruct-Algorithms/Data-structures/loopqueue"
	. "Go-DataStruct-Algorithms/Data-structures/linkedlist"
	"Go-DataStruct-Algorithms/util"
)

func Test_QueueBanch(t *testing.T)  {
	var queue Queue
	opCount := 100000

	queue = CreatedefaultQueue(reflect.Int)
	latency := CompareQueueTime(queue,opCount)
	t.Logf("ArrayQueue, time is %-13s",latency)

	queue = CreatedefaultLoopQueue(reflect.Int)
	latency = CompareQueueTime(queue,opCount)
	t.Logf("LoopQueue, time is %-13s",latency)

	queue = CreateLinkedListQueue()
	latency = CompareQueueTime(queue,opCount)
	t.Logf("LinkedListQueue, time is %-13s",latency)
}

func CompareQueueTime(queue Queue,opCount int) time.Duration {
	start := time.Now().UTC()

	randNum := util.Rand_generator()
	for i := 0; i < opCount ; i++  {
		queue.Enqueue(<-randNum)
	}
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}

	end := time.Now().UTC()
	latency := end.Sub(start)

	return latency
}

