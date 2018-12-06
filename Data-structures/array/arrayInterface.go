package array

type ArrayList interface {
	Add(int, interface{})
	AddFirst(interface{})
	AddLast(interface{})

	IsEmpty() bool
	Contains(interface{}) bool

	Find(interface{}) int
	FindAll(interface{}) []int

	Set(int, interface{})
	Get(int) interface{}
	GetArrCap() int
	GetArrLen() int

	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	RemoveAllElement(interface{}) bool
	RemoveElement(interface{}) bool

	String() string
}
