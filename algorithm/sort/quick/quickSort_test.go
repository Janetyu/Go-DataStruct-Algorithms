package quick

import "testing"

func Test_QuickSort1(t *testing.T) {
	arr := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	arr = QuickSort1(arr)
	t.Log(arr)
}

func Test_QuickSort2(t *testing.T) {
	arr := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	arr = QuickSort2(arr)
	t.Log(arr)
}

func Test_QuickSort3(t *testing.T) {
	arr := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	arr = QuickSort2(arr)
	t.Log(arr)
}
