package shell

// 希尔排序（缩小增量排序），是简单插入算法的增强版

func ShellSort(arr []int) []int {
	var temp int
	len := len(arr) // 10
	gap := 1
	for gap < len / 3 { // 动态定义间隔序列
		gap = gap * 3 + 1 // grap==4
	}

	for ;gap > 0; gap = gap / 3 {
		for i := gap; i < len; i++ {
			temp = arr[i]
			var j int
			for j = i - gap;j >= 0 && arr[j] > temp;j -= gap {
				arr[j + gap] = arr[j]
			}
			arr[j + gap] = temp
		}
	}

	return arr
}