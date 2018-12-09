package stack

import (
	"testing"
	"reflect"
	"fmt"
)

func Test_Stack(t *testing.T) {
	stack := CreatedefaultStack(reflect.Int)
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
