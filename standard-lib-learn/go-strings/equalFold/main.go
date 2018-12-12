package main

/*
EqualFold 判断 s 和 t 是否相等。忽略大小写，同时它还会对特殊字符进行转换
比如将“ϕ”转换为“Φ”、将“Ǆ”转换为“ǅ”等，然后再进行比较

func EqualFold(s, t string) bool
*/

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello 世界! ϕ Ǆ"
	s2 := "hello 世界! Φ ǅ"
	b := strings.EqualFold(s1, s2)
	fmt.Printf("%v\n", b) // true
}
