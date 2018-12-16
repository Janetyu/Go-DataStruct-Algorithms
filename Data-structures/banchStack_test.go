package Data_structures

import (
	"testing"
	"reflect"
	"time"

	. "Go-DataStruct-Algorithms/Data-structures/stack"
	. "Go-DataStruct-Algorithms/Data-structures/linkedlist"

	"Go-DataStruct-Algorithms/util"
)

func Test_StackBanch(t *testing.T)  {
	var stack Stack
	opCount := 100000

	stack = CreatedefaultStack(reflect.Int)
	latency := CompareStackTime(stack,opCount)
	t.Logf("ArrayStack, time is %-13s",latency)

	stack = CreateLinkedListStack()
	latency = CompareStackTime(stack,opCount)
	t.Logf("LinkedListStack, time is %-13s",latency)

	// 其实这个时间比较很复杂，因为LinkedListStack中包含更多的new操作

}

func CompareStackTime(stack Stack,opCount int) time.Duration {
	start := time.Now().UTC()

	randNum := util.Rand_generator()
	for i := 0; i < opCount ; i++  {
		stack.Push(<-randNum)
	}
	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	end := time.Now().UTC()
	latency := end.Sub(start)

	return latency
}
