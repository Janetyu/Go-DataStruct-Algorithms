package loopqueue

import "reflect"

type LoopQueue struct {
	arrType     reflect.Kind
	data        []interface{}
	front, tail int
	size        int
}
