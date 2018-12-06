package array

import (
	"bytes"
	"fmt"
	"reflect"
)

type Array struct {
	arrType reflect.Kind
	data    []interface{}
	size    int
}

// 指定容量创建数组，可自定义数组类型
func CreateArray(arrType reflect.Kind, cap int) *Array {
	data := make([]interface{}, cap) // 若不指定 cap 的大小，则自动设置为 len == cap
	size := 0
	return &Array{arrType, data, size}
}

// 无参数创建默认容量为10的数组，可自定义数组类型
func CreatedefaultArray(arrType reflect.Kind) *Array {
	data := make([]interface{}, 10)
	size := 0
	return &Array{arrType, data, size}
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
func (arr *Array) AddLast(e interface{}) {
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
func (arr *Array) AddFirst(e interface{}) {
	arr.Add(0, e)
}

// 向数组中插入一个元素
func (arr *Array) Add(index int, e interface{}) {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("Add failed.element type worried.")
	}

	if arr.size == cap(arr.data) {
		//panic("Add failed.Array is full.")
		arr = resize(2*cap(arr.data), arr)
		fmt.Println(arr.String())
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
func (arr *Array) Get(index int) interface{} {
	if index < 0 || index >= arr.size {
		panic("Get failed.Index is illegal.")
	}
	return arr.data[index]
}

// 修改index索引位置的元素
func (arr *Array) Set(index int, e interface{}) {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("Set failed.element type worried.")
	}

	if index < 0 || index >= arr.size {
		panic("Set failed.Index is illegal.")
	}
	arr.data[index] = e
}

// 查找数组中是否有元素e
func (arr *Array) Contains(e interface{}) bool {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("IsContains failed.element type worried.")
	}

	for _, v := range arr.data {
		if e == v {
			return true
		}
	}
	return false
}

// 查找数组中元素e所在的索引，如果不存在元素e，则返回-1
func (arr *Array) Find(e interface{}) int {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("Find failed.element type worried.")
	}

	for i, v := range arr.data {
		if e == v {
			return i
		}
	}
	return -1
}

// 查找数组中所有元素e所在的索引，如果不存在元素e，则返回空slice
func (arr *Array) FindAll(e interface{}) []int {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("FindAll failed.element type worried.")
	}

	indexs := []int{}
	for i, v := range arr.data {
		if e == v {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

// 从数组中删除index位置的元素，返回删除的元素
func (arr *Array) Remove(index int) interface{} {
	if index < 0 || index >= arr.size {
		panic("Remove failed.Index is illegal.")
	}
	res := arr.data[index]
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	//arr.data[arr.size] = 0 // loitering objects != memory leak

	if arr.size == cap(arr.data)/2 {
		arr = resize(cap(arr.data)/2, arr)
	}

	return res
}

// 从数组中删除第一个元素，返回删除的元素
func (arr *Array) RemoveFirst() interface{} {
	return arr.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (arr *Array) RemoveLast() interface{} {
	return arr.Remove(arr.size - 1)
}

// 从数组中删除元素e
func (arr *Array) RemoveElement(e interface{}) bool {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("RemoveElement failed.element type worried.")
	}

	index := arr.Find(e)
	if index != -1 {
		arr.Remove(index)
		return true
	}
	return false
}

// 从数组中删除所有的元素e
func (arr *Array) RemoveAllElement(e interface{}) bool {
	etype := reflect.TypeOf(e).Kind()
	if etype != arr.arrType {
		panic("RemoveAllElement failed.element type worried.")
	}

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
		buffer.WriteString(fmt.Sprintf("%v", arr.data[i]))
		if i != arr.size-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(str3)
	return buffer.String()
}

// 实现动态数组
func resize(cap int, arr *Array) *Array {
	newdata := make([]interface{}, cap)
	for i := 0; i < arr.size; i++ {
		newdata[i] = arr.data[i]
	}
	arr.data = newdata
	return arr
}
