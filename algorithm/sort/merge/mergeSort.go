package merge

// 归并排序

func MergeSort(arr []int) []int { // 采用自上而下的递归方法
	len := len(arr)
	if len < 2 {
		return arr
	}

	middle := len / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	result := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			l, leftfirst := removefirst(left)
			left = l
			result = append(result, leftfirst)
		} else {
			r, rightfirst := removefirst(right)
			right = r
			result = append(result, rightfirst)
		}
	}

	for len(left) > 0 {
		l, leftfirst := removefirst(left)
		left = l
		result = append(result, leftfirst)
	}

	for len(right) > 0 {
		r, rightfirst := removefirst(right)
		right = r
		result = append(result, rightfirst)
	}

	return result
}

func removefirst(arr []int) ([]int, int) {
	temp := arr[0]
	arr = arr[1:]
	return arr, temp
}
