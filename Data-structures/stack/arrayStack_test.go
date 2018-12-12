package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := CreatedefaultStack(reflect.Int)
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)

	str1 := "()[]{}"
	str2 := "([)]"
	t.Logf("the string is %s, and isVaild result is %v", str1, isValid(str1))
	t.Logf("the string is %s, and isVaild result is %v", str2, isValid(str2))
}
