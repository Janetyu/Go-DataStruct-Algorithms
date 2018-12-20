package recursive

func Sum(arr []int) int {
	return	sum(arr, 0)
}

// 计算 arr[l...n) 范围里的数字和
func sum(arr []int, l int) int {
	if l == len(arr) {
		return 0
	}

	return arr[l] + sum(arr,l + 1)
}
