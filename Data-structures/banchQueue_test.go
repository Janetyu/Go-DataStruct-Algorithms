package Data_structures

import (
	"testing"
	"reflect"
	"time"
	"math/rand"
	"math"

	. "Go-DataStruct-Algorithms/Data-structures/queue"
	. "Go-DataStruct-Algorithms/Data-structures/loopqueue"
)

func Test_QueueBanch(t *testing.T)  {
	var queue Queue
	opCount := 10000

	queue = CreatedefaultQueue(reflect.Int)
	latency := CompareTime(queue,opCount)
	t.Logf("ArrayQueue, time is %-13s",latency)

	queue = CreatedefaultLoopQueue(reflect.Int)
	latency = CompareTime(queue,opCount)
	t.Logf("LoopQueue, time is %-13s",latency)

}

func CompareTime(queue Queue,opCount int) time.Duration {
	start := time.Now().UTC()

	randNum := rand_generator()
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

func rand_generator() chan int{
	out:=make(chan int)
	go func(){
		for{
			rand.Seed(time.Now().Unix())
			out <- rand.Intn(math.MaxInt64)

		}
	}()
	return out
}
