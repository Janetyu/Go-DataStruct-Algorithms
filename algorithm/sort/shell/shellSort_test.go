package shell

import "testing"

func Test_ShellSort(t *testing.T) {
	arr := []int{84,83,88,87,61,50,70,60,80,99}
	arr = ShellSort(arr)
	t.Log(arr)
}

