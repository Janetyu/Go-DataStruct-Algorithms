package quick

func QuickSort1(arr []int) []int {
	left := 0
	right := len(arr)
	len := len(arr)
	if left < right {
		key := arr[left]
		i := left
		j := right - 1

		for {
			for i+1 < len {
				i++
				if key >= arr[j] {
					break
				}
				j--
			}

			if i >= j {
				break
			}

			arr = swap(arr, i, j)
		}
	}

	return arr
}

func QuickSort2(arr []int) []int {
	return nil
}

func QuickSort3(arr []int) []int {
	return nil
}

func swap(arr []int, a, b int) []int {
	arr[a], arr[b] = arr[b], arr[a]
	return arr
}
