package array

type Array struct {
	data []int
	size int
}

// 指定容量创建数组
func (arr *Array)CreateArray(cap int)  {
	arr.data = make([]int,0,cap)
	arr.size = 0
}

// 无参数创建默认容量为10的数组
func (arr *Array)CreatedefaultArray() {
	arr.data = make([]int,0,10)
	arr.size = 0
}

// 获取数组的长度
func (arr *Array)GetArrLen() int {
	return arr.size
}

// 获取数组的容量
func (arr *Array)GetArrCap() int {
	return cap(arr.data)
}

// 判断数组是否为空
func (arr *Array)IsEmpty() bool {
	return arr.size == 0
}

// 向数组的末尾添加一个元素
func (arr *Array)AddLast(e int) {

	//if arr.size == cap(arr.data) {
	//	panic("AddLast failed.Array is full.")
	//}
	//
	//arr.data[arr.size] = e
	//
	//arr.size += 1

	arr.Add(arr.size,e)
}

// 向数组的首部添加一个元素
func (arr *Array)AddFirst(e int)  {
	arr.Add(0,e)
}

// 向数组中插入一个元素
func (arr *Array)Add(index,e int)  {

	if arr.size == cap(arr.data) {
		panic("Add failed.Array is full.")
	}

	if index <= 0 || index >= cap(arr.data) {
		panic("Add failed.Require index >= 0 and <= arr.capacity")
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i + 1] = arr.data[i]
	}

	arr.data[index] = e
	arr.size += 1
}


