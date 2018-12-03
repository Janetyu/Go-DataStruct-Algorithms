package main

/*
FieldsFunc 以一个或多个满足 f(rune) 的字符为分隔符，
将 s 切分成多个子串，结果中不包含分隔符本身。
如果 s 中没有满足 f(rune) 的字符，则返回一个空列表。

func FieldsFunc(s string, f func(rune) bool) []string
 */

import (
	"strings"
	"fmt"
)

func isSlash(r rune) bool {
	return r == '\\' || r == '/'
}

func main() {
	s := "C:\\Windows\\System32\\FileName"
	ss := strings.FieldsFunc(s, isSlash)
	fmt.Printf("%q\n", ss) // ["C:" "Windows" "System32" "FileName"]
}
