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
}
