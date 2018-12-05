package array

import (
	"testing"
)

func Test_Array(t *testing.T)  {
	arr := CreateArray(20)
	for i := 0; i < 10 ; i++ {
		arr.AddLast(i)
	}
	t.Log(arr.String())
}