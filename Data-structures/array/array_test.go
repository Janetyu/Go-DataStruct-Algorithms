package array

import (
	"reflect"
	"testing"
)

func Test_Array(t *testing.T) {
	arr := CreateArray(reflect.Int, 20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	t.Log(arr.String())

	arr.Add(0, 100)
	t.Log(arr.String())

	arr.AddFirst(200)
	t.Log(arr.String())

	arr.Set(0, -1)
	t.Log(arr.String())

	ok := arr.RemoveAllElement(-2)
	t.Log("Remove all: ", ok)
	t.Log(arr.String())

	ok = arr.RemoveAllElement(-1)
	t.Log("Remove all: ", ok)
	t.Log(arr.String())

	arr.Remove(0)
	t.Log(arr.String())

	arr.RemoveElement(3)
	t.Log(arr.String())

	var arrList ArrayList // 声明接口
	arrList = CreatedefaultArray(reflect.String)
	arrList.AddFirst("first")
	arrList.AddLast("two")
	arrList.AddLast("three")
	//arrList.RemoveElement(1) // RemoveElement failed.element type worried.
	arrList.RemoveElement("three")
	t.Log(arrList.String())
}
