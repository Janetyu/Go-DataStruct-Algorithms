package array

import (
	"testing"
)

func Test_Array(t *testing.T) {
	arr := CreateArray(20)
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
}
