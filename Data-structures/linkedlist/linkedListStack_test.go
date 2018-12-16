package linkedlist

import (
	"testing"
	"fmt"
)

func Test_LinkedListStack(t *testing.T)  {
	stack := CreateLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack.String())
	}

	stack.Pop()
	fmt.Println(stack.String())
}
