package array

import (
	"bytes"
	"fmt"
	"strconv"
)

type Array struct {
	data []int
	size int
}

// 指定容量创建数组
func CreateArray(cap int) *Array {
	data := make([]int,cap) // 若不指定 cap 的大小，则自动设置为 len == cap
	size := 0
	return &Array{data,size}
}

// 无参数创建默认容量为10的数组
func CreatedefaultArray() *Array {
	data := make([]int,10)
	size := 0
	return &Array{data,size}
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

	if index < 0 || index > arr.size {
		panic("Add failed.Require index >= 0 and index <= size")
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i + 1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size += 1
}

func (arr *Array)String() string {
	var buffer bytes.Buffer
	str1 := fmt.Sprintf("Array: size = %d, capacity = %d \n",arr.size, cap(arr.data))
	str2 := "["
	str3 := "]"
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	for i := 0; i < arr.size ; i++ {
		buffer.WriteString(strconv.Itoa(arr.data[i]))
		if i != arr.size - 1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(str3)
	return buffer.String()
}


