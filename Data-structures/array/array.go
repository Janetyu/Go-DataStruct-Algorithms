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
	data := make([]int, cap) // 若不指定 cap 的大小，则自动设置为 len == cap
	size := 0
	return &Array{data, size}
}

// 无参数创建默认容量为10的数组
func CreatedefaultArray() *Array {
	data := make([]int, 10)
	size := 0
	return &Array{data, size}
}

// 获取数组的长度
func (arr *Array) GetArrLen() int {
	return arr.size
}

// 获取数组的容量
func (arr *Array) GetArrCap() int {
	return cap(arr.data)
}

// 判断数组是否为空
func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}

// 向数组的末尾添加一个元素
func (arr *Array) AddLast(e int) {
	//if arr.size == cap(arr.data) {
	//	panic("AddLast failed.Array is full.")
	//}
	//
	//arr.data[arr.size] = e
	//
	//arr.size += 1

	arr.Add(arr.size, e)
}

// 向数组的首部添加一个元素
func (arr *Array) AddFirst(e int) {
	arr.Add(0, e)
}

// 向数组中插入一个元素
func (arr *Array) Add(index, e int) {

	if arr.size == cap(arr.data) {
		panic("Add failed.Array is full.")
	}

	if index < 0 || index > arr.size {
		panic("Add failed.Require index >= 0 and index <= size")
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size += 1
}

// 获取index索引位置的元素
func (arr *Array) Get(index int) int {
	if index < 0 || index >= arr.size {
		panic("Get failed.Index is illegal.")
	}
	return arr.data[index]
}

// 修改index索引位置的元素
func (arr *Array) Set(index, e int) {
	if index < 0 || index >= arr.size {
		panic("Set failed.Index is illegal.")
	}
	arr.data[index] = e
}

// 查找数组中是否有元素e
func (arr *Array) Contains(e int) bool {
	for _, v := range arr.data {
		if e == v {
			return true
		}
	}
	return false
}

// 查找数组中元素e所在的索引，如果不存在元素e，则返回-1
func (arr *Array) Find(e int) int {
	for i, v := range arr.data {
		if e == v {
			return i
		}
	}
	return -1
}

// 查找数组中所有元素e所在的索引，如果不存在元素e，则返回空slice
func (arr *Array) FindAll(e int) []int {
	indexs := []int{}
	for i, v := range arr.data {
		if e == v {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

// 从数组中删除index位置的元素，返回删除的元素
func (arr *Array) Remove(index int) int {
	if index < 0 || index >= arr.size {
		panic("Remove failed.Index is illegal.")
	}
	res := arr.data[index]
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	return res
}

// 从数组中删除第一个元素，返回删除的元素
func (arr *Array) RemoveFirst() int {
	return arr.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (arr *Array) RemoveLast() int {
	return arr.Remove(arr.size - 1)
}

// 从数组中删除元素e
func (arr *Array) RemoveElement(e int) bool {
	index := arr.Find(e)
	if index != -1 {
		arr.Remove(index)
		return true
	}
	return false
}

// 从数组中删除所有的元素e
func (arr *Array) RemoveAllElement(e int) bool {
	indexs := arr.FindAll(e)

	if len(indexs) == 0 {
		//panic("Remove failed.Without the value in Array.")
		return false
	}

	for _, v := range indexs {
		arr.Remove(v)
	}

	return true

}

func (arr *Array) String() string {
	var buffer bytes.Buffer
	str1 := fmt.Sprintf("Array: size = %d, capacity = %d \n", arr.size, cap(arr.data))
	str2 := "["
	str3 := "]"
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	for i := 0; i < arr.size; i++ {
		buffer.WriteString(strconv.Itoa(arr.data[i]))
		if i != arr.size-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(str3)
	return buffer.String()
}
