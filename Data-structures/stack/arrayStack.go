package stack

import (
	"Go-DataStruct-Algorithms/Data-structures/array"
	"bytes"
	"fmt"
	"reflect"
)

type ArrayStack struct {
	array *array.Array
}

// 指定容量创建栈，可自定义数组类型
func CreateStack(stackType reflect.Kind, cap int) *ArrayStack {
	return &ArrayStack{array.CreateArray(stackType, cap)}
}

// 无参数创建默认容量为10的栈，可自定义数组类型
func CreatedefaultStack(stackType reflect.Kind) *ArrayStack {
	return &ArrayStack{array.CreatedefaultArray(stackType)}
}

func (stack *ArrayStack) GetSize() int {
	return stack.array.GetArrLen()
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.array.IsEmpty()
}

func (stack *ArrayStack) GetCapacity() int {
	return stack.array.GetArrCap()
}

func (stack *ArrayStack) Push(e interface{}) {
	stack.array.AddLast(e)
}

func (stack *ArrayStack) Pop() interface{} {
	return stack.array.RemoveLast()
}

func (stack *ArrayStack) Peek() interface{} {
	return stack.array.GetLast()
}

func (stack *ArrayStack) String() string {
	var buffer bytes.Buffer
	str1 := fmt.Sprintf("Stack: size = %d, capacity = %d \n", stack.array.GetArrLen(), stack.array.GetArrCap())
	str2 := "["
	str3 := "] top"
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	for i := 0; i < stack.array.GetArrLen(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", stack.array.Get(i)))
		if i != stack.array.GetArrLen()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(str3)
	return buffer.String()
}
