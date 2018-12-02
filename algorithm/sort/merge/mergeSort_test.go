package merge

import "testing"

func Test_MergeSort(t *testing.T) {
	arr := []int{84,83,88,87,61,50,70,60,80,99}
	arr = MergeSort(arr)
	t.Log(arr)
}
